package ordermanager

import (
	"MiSwapSync/stores/gdb/multi"
	"MiSwapSync/stores/xkv"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RussellLuo/timingwheel"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
	"log/slog"
	"sync/atomic"
	"time"
)

const (
	// WheelSize : length of time wheel, 1s for a unit, 1h for a cycle
	WheelSize           = 3600
	List                = 3
	CacheOrdersQueuePre = "cache:ms:orders:%s"
)

type Order struct {
	// order Id
	orderID        string
	CollectionAddr string
	// chain suffix name
	ChainSuffix string
	// expireIn - createIn (unit: s)
	CycleCount int64
	// position of the task on the time wheel
	WheelPosition int64
	Next          *Order
}

type wheel struct {
	// linked list
	NotifyActivities *Order
}

//	type OrderManager struct {
//		chain        string
//		TimeWheel    [WheelSize]wheel // 时间轮
//		CurrentIndex int64            // current time wheel index
//
//		//collectionOrders map[string]*collectionTradeInfo
//
//		collectionListedCh chan string
//		project            string
//
//		Xkv *xkv.Store
//		DB  *gorm.DB
//		Mux *sync.RWMutex
//	}
type collectionTradeInfo struct {
	floorPrice decimal.Decimal
	orders     *PriorityQueueMap // 优先级队列
}
type ListingInfo struct {
	ExpireIn       int64           `json:"expire_in"`
	OrderId        string          `json:"order_id"`
	CollectionAddr string          `json:"collection_addr"`
	TokenID        string          `json:"token_id"`
	Price          decimal.Decimal `json:"price"`
	Maker          string          `json:"maker"`
}

type OrderManager struct {
	DB                 *gorm.DB
	Xkv                *xkv.Store
	project            string
	chain              string
	collectionListedCh chan string
	timingWheel        *timingwheel.TimingWheel
	taskCount          atomic.Int64 // 追踪时间轮中待处理任务总数
	collectionOrders   map[string]*collectionTradeInfo
}

// NewOrderManager 创建 OrderManager 实例并启动时间轮
func NewOrderManager(
	db *gorm.DB,
	project, chain string,
) *OrderManager {
	// 创建时间轮：tick=1秒，wheelSize=60（一圈覆盖60秒）
	// 库内部使用层级时间轮（Hierarchical Timing Wheels），
	// 超出单圈时长的延迟会自动提升到上层时间轮，无需手动处理 CycleCount
	tw := timingwheel.NewTimingWheel(time.Second, 4096)
	tw.Start()

	return &OrderManager{
		DB:          db,
		project:     project,
		chain:       chain,
		timingWheel: tw,
	}
}

func (om *OrderManager) Start(ctx context.Context) {
	threading.GoSafe(func() { om.ListenNewListingLoop(ctx) }) // 处理新订单
	threading.GoSafe(func() { om.orderExpiryProcess(ctx) })   // 处理订单过期状态
	threading.GoSafe(func() { om.floorPriceProcess(ctx) })    // 处理floor price更新
	threading.GoSafe(func() { om.listCountProcess(ctx) })     // 处理listCount更新
}

func (om *OrderManager) ListenNewListingLoop(ctx context.Context) {
	key := fmt.Sprintf(CacheOrdersQueuePre, om.chain)
	for {
		select {
		case <-ctx.Done():
			slog.InfoContext(ctx, "ListenNewListingLoop stopped")
			return
		default:
		}
		result, err := om.Xkv.Lpop(key)
		if err != nil || result == "" {
			if err != nil && !errors.Is(err, redis.Nil) {
				slog.WarnContext(context.Background(), "failed to get order from cache", "error", err, "result", result)
			}
			//防止空转，降低资源消耗.  用 select+timer 替代 time.Sleep，使等待期间也能响应 ctx
			select {
			case <-ctx.Done():
				slog.InfoContext(ctx, "ListenNewListingLoop stopped during backoff")
				return
			case <-time.After(1 * time.Second):
			}
			//time.Sleep(1 * time.Second)
			continue
		}

		slog.InfoContext(context.Background(), "get listing from cache", "result", result)

		var listing ListingInfo
		if err := json.Unmarshal([]byte(result), &listing); err != nil {
			slog.WarnContext(context.Background(), "failed to unmarshal order info", "error", err)
			continue
		}

		if listing.OrderId == "" {
			slog.WarnContext(context.Background(), "invalid order id")
			continue
		}

		if listing.ExpireIn < time.Now().Unix() { // 订单已经过期
			slog.WarnContext(context.Background(), "expired activity order", "order_id", listing.OrderId)
			// 更新订单状态
			if err := om.updateOrdersStatus(listing.OrderId, multi.OrderStatusExpired); err != nil {
				slog.Error("failed to expire activity order", "error", err)
			}
			// 添加更新floor price事件
			if err := om.addUpdateFloorPriceEvent(&TradeEvent{
				EventType:      Expired,
				CollectionAddr: listing.CollectionAddr,
				TokenID:        listing.TokenID,
				OrderId:        listing.OrderId,
				From:           listing.Maker,
			}); err != nil {
				slog.Error("failed to enqueue floor price recalculation for expired order", "order_id", listing.OrderId, "error", err)
			}
			continue
		} else { // 订单未过期
			if err := om.addUpdateFloorPriceEvent(&TradeEvent{
				EventType:      Listing,
				CollectionAddr: listing.CollectionAddr,
				TokenID:        listing.TokenID,
				OrderId:        listing.OrderId,
				Price:          listing.Price,
				From:           listing.Maker,
			}); err != nil {
				slog.Error("failed to enqueue floor price recalculation for new listing", "order_id", listing.OrderId, "price", listing.Price, "collection_addr", listing.CollectionAddr)
			}
			// 添加到订单过期检查队列
			delaySeconds := listing.ExpireIn - time.Now().Unix()
			om.addToOrderExpiryCheckQueue(time.Duration(delaySeconds)*time.Second, om.chain, listing.OrderId, listing.CollectionAddr)
		}
	}
}
