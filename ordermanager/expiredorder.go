package ordermanager

import (
	"MiSwapSync/stores/gdb"
	"MiSwapSync/stores/gdb/multi"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"log/slog"
	"time"
)

const (
	Buy EventType = iota + 1
	Mint
	Listing
	Cancel
	Transfer         EventType = 8
	Expired          EventType = 9
	ImportCollection EventType = 10
	UpdateCollection EventType = 11
)

type EventType int
type TradeEvent struct {
	EventType      EventType       `json:"event_type"`
	CollectionAddr string          `json:"collection_addr"`
	TokenID        string          `json:"token_id"`
	OrderId        string          `json:"order_id"`
	OrderHash      string          `json:"order_hash"`
	Price          decimal.Decimal `json:"price"`
	From           string          `json:"from"`
	To             string          `json:"to"`
	TxHash         string          `json:"txHash"`
}

// ============================================================
// orderExpiryProcess - 订单过期处理主入口
// ============================================================

// OrderExpiryProcess 启动订单过期处理流程。
func (om *OrderManager) orderExpiryProcess(ctx context.Context) {
	// 1. 使用 defer recover 来捕获可能的 panic，防止主协程死掉
	defer func() {
		if r := recover(); r != nil {
			slog.Warn("[Order Manage] dq process recovered: " + fmt.Sprintf("%v", r))
		}
	}()

	// 2. 启动时从数据库加载所有活跃订单到时间轮队列中
	if err := om.loadOrdersToQueue(); err != nil {
		slog.Warn("[Order Manage] load orders to queue", zap.Error(err))
		return
	}

	// 3. 时间轮库内部已启动独立的 goroutine 进行任务调度，
	//    我们只需保持主协程存活即可。
	// ctx 被 cancel 时，ctx.Done() 自动解除阻塞
	<-ctx.Done()
	slog.Info("[Order Manage] order expiry process stopping, shutting down time wheel")
	om.timingWheel.Stop()
}

// ============================================================
// loadOrdersToQueue - 启动时加载订单
// ============================================================

// loadOrdersToQueue 函数负责在系统启动时加载所有活跃订单并处理它们的过期状态。
// 主要功能包括:
// 1. 从数据库分批加载所有活跃订单
// 2. 检查每个订单是否已过期:
//   - 已过期的订单: 更新状态为过期并触发地板价更新事件
//   - 未过期的订单: 添加到时间轮等待过期检查
func (om *OrderManager) loadOrdersToQueue() error {
	// 分批加载所有活跃订单
	var totalOrders []*multi.Order
	var id int64
	for {
		var orders []*multi.Order
		if err := om.DB.Table(gdb.GetMultiProjectOrderTableName(om.project, om.chain)).
			Select("id, order_id, collection_address, price, expire_time").
			Where("order_status = ? and id > ?", multi.OrderStatusActive, id).
			Order("id asc").Limit(1000).
			Scan(&orders).Error; err != nil {
			return errors.Wrap(err, "failed to get collection orders")
		}
		totalOrders = append(totalOrders, orders...)
		if len(orders) < 1000 {
			break
		}
		id = orders[len(orders)-1].ID
	}

	if totalOrders == nil || len(totalOrders) == 0 {
		return nil
	}

	// 分别记录已过期和未过期的订单
	var expiredOrderIDs []int64
	var expiredOrders []*multi.Order
	for _, order := range totalOrders {
		if order.ExpireTime < time.Now().Unix() { // 已过期
			expiredOrderIDs = append(expiredOrderIDs, order.ID)
			expiredOrders = append(expiredOrders, order)
		} else { // 未过期，添加到时间轮
			delaySeconds := order.ExpireTime - time.Now().Unix()
			om.addToOrderExpiryCheckQueue(time.Duration(delaySeconds)*time.Second, om.chain, order.OrderID, order.CollectionAddress)
		}
	}

	// 批量更新已过期订单的状态
	for i := 0; i < len(expiredOrderIDs); i += MaxBatchReqNum {
		end := i + MaxBatchReqNum
		if i+MaxBatchReqNum >= len(expiredOrderIDs) {
			end = len(expiredOrderIDs)
		}
		if err := om.DB.Table(gdb.GetMultiProjectOrderTableName(om.project, om.chain)).
			Where("id in (?)", expiredOrderIDs[i:end]).Update("order_status", multi.OrderStatusExpired).Error; err != nil {
			return errors.Wrap(err, "failed to update expired orders status")
		}
	}

	// 为每个过期订单触发地板价更新事件
	for _, order := range expiredOrders {
		if err := om.addUpdateFloorPriceEvent(&TradeEvent{
			EventType:      Expired,
			CollectionAddr: order.CollectionAddress,
			OrderId:        order.OrderID,
		}); err != nil {
			return errors.Wrap(err, "failed to add update floor price event")
		}
	}

	return nil
}

// ============================================================
// addToOrderExpiryCheckQueue - 添加订单到时间轮
// ============================================================

// addToOrderExpiryCheckQueue 函数用于将订单添加到时间轮中进行过期检查。
//
//		现在只需调用 AfterFunc 指定延迟时间即可，库内部自动处理所有调度逻辑
//	    （包括跨圈调度、溢出到上层时间轮等）
//
// 参数说明:
// - delay: 延迟时间，即订单过期前的剩余时间
// - chainSuffix: 链标识
// - orderId: 订单ID
// - collectionAddr: NFT集合地址
func (om *OrderManager) addToOrderExpiryCheckQueue(delay time.Duration, chainSuffix string, orderId string, collectionAddr string) {
	// timingwheel 库内部已保证线程安全，无需手动加锁
	// 使用 AfterFunc 添加一个一次性定时任务
	om.timingWheel.AfterFunc(delay, func() {
		// 异步更新订单状态，避免阻塞时间轮内部协程
		go func() {
			if err := om.updateOrderState(orderId, collectionAddr); err != nil {
				slog.Error("failed to update order status",
					zap.Error(err),
					zap.String("chain", chainSuffix),
					zap.String("order_id", orderId))
			}
		}()
	})
}

// ============================================================
// GetWheelTaskQuantity - 获取时间轮任务数量
// ============================================================

// GetWheelTaskQuantity 返回时间轮中待处理任务的大致数量。
// 改写说明：
// - 原实现通过遍历时间槽链表获取特定槽位的任务数
// - timingwheel 库不暴露底层槽位结构，因此改为返回全局待处理任务总数
// - 通过 atomic.Int64 在添加/执行任务时增减计数来维护
func (om *OrderManager) GetWheelTaskQuantity() int64 {
	return om.taskCount.Load()
}

// ============================================================
// updateOrderState - 订单状态更新处理
// ============================================================

// updateOrderState : handler function by sql transaction
func (om *OrderManager) updateOrderState(orderId string, collectionAddr string) error {
	// update orders status to expire
	if err := om.updateOrdersStatus(orderId, multi.OrderStatusExpired); err != nil {
		return errors.Wrap(err, "failed to expire order")
	}

	// update floor price
	if err := om.addUpdateFloorPriceEvent(&TradeEvent{
		EventType:      Expired,
		OrderId:        orderId,
		CollectionAddr: collectionAddr,
	}); err != nil {
		return errors.Wrap(err, "failed to add update floor price event")
	}

	return nil
}

// updateOrdersStatus 更新订单状态
func (om *OrderManager) updateOrdersStatus(orderID string, orderStatus int) error {
	if err := om.DB.Table(gdb.GetMultiProjectOrderTableName(om.project, om.chain)).
		Where("order_id = ?", orderID).Update("order_status", orderStatus).Error; err != nil {
		return errors.Wrap(err, "failed to update expired orders status")
	}
	return nil
}

func (om *OrderManager) addUpdateFloorPriceEvent(event *TradeEvent) error {
	// 验证事件合法性
	// 如果是Transfer/ImportCollection/UpdateCollection事件,必须有OrderId
	if event.EventType != Transfer && event.EventType != ImportCollection && event.EventType != UpdateCollection && event.OrderId == "" {
		return errors.New("invalid update collection floor price. event order id is null")
	}
	// 如果是Listing事件,价格不能为0
	if event.EventType == Listing && event.Price.IsZero() {
		return errors.New("invalid update collection floor price. price is 0")
	}

	// 序列化事件
	rawEvent, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "failed to marshal event")
	}

	// 获取Redis队列key
	key := fmt.Sprintf("cache:trade:events:%s", om.chain)
	// 将事件添加到Redis队列
	if _, err := om.Xkv.Rpush(key, string(rawEvent)); err != nil {
		return errors.Wrap(err, "failed to push trade event to queue")
	}
	return nil
}
