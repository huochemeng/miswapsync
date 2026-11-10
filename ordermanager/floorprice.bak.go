package ordermanager

//
//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"log/slog"
//	"strings"
//	"time"
//
//	"github.com/redis/go-redis/v9"
//	"github.com/shopspring/decimal"
//)
//
//// EventHandler 定义事件处理策略接口
//type EventHandler func(ctx context.Context, om *OrderManager, event TradeEvent) error
//
//// OrderManager 订单管理器
//
//// buildHandlerMap 使用表驱动法注册事件处理器
//func (om *OrderManager) buildHandlerMap() map[EventType]EventHandler {
//	return map[EventType]EventHandler{
//		Listing:          om.handleListing,
//		Cancel:           om.handleCancelOrExpired,
//		Expired:          om.handleCancelOrExpired,
//		Buy:              om.handleBuyOrTransfer,
//		Transfer:         om.handleBuyOrTransfer,
//		ImportCollection: om.handleImportCollection,
//		UpdateCollection: om.handleUpdateCollection,
//	}
//}
//
//// FloorPriceProcess 主入口：ctx 取消时优雅退出监听循环
//func (om *OrderManager) FloorPriceProcess(ctx context.Context) {
//	key := genTradeEventsCacheKey(om.chain)
//	om.clearRemainingEvents(ctx, key)
//
//	if err := om.loadCollectionTradeInfo(); err != nil {
//		om.logger.ErrorContext(ctx, "[Order Manage] load orders to queue failed", "error", err)
//		return
//	}
//
//	handlers := om.buildHandlerMap()
//
//	for {
//		// ✅ 核心改动：每次循环检查 ctx 是否已取消
//		select {
//		case <-ctx.Done():
//			om.logger.InfoContext(ctx, "floor price process stopped", "reason", ctx.Err())
//			return
//		default:
//		}
//
//		event, ok := om.fetchNextEvent(ctx, key)
//		if !ok {
//			// 使用带 ctx 的 sleep，ctx 取消时立即返回而非等满 1s
//			select {
//			case <-ctx.Done():
//				om.logger.InfoContext(ctx, "floor price process stopped during backoff", "reason", ctx.Err())
//				return
//			case <-time.After(1 * time.Second):
//			}
//			continue
//		}
//
//		handler, exists := handlers[event.EventType]
//		if !exists {
//			om.logger.ErrorContext(ctx, "unsupported event type",
//				"event_type", int(event.EventType),
//				"raw_event", event.rawJSON,
//			)
//			continue
//		}
//
//		if err := handler(ctx, om, event); err != nil {
//			om.logger.ErrorContext(ctx, "handle trade event failed",
//				"event_type", int(event.EventType),
//				"collection_addr", event.CollectionAddr,
//				"order_id", event.OrderId,
//				"error", err,
//			)
//		}
//	}
//}
//
//// fetchNextEvent 封装 Redis LPOP + JSON 反序列化 + 基础校验
//func (om *OrderManager) fetchNextEvent(ctx context.Context, key string) (TradeEvent, bool) {
//	result, err := om.Xkv.Lpop(key) // ✅ Lpop 接受 ctx，支持超时/取消
//	if err != nil || result == "" {
//		if err != nil && !errors.Is(err, redis.Nil) {
//			om.logger.WarnContext(ctx, "failed on get trade events from cache", "error", err)
//		}
//		return TradeEvent{}, false
//	}
//
//	var event TradeEvent
//	if err := json.Unmarshal([]byte(result), &event); err != nil {
//		om.logger.WarnContext(ctx, "failed on unmarshal trade event",
//			"error", err,
//			"raw", result,
//		)
//		return TradeEvent{}, false
//	}
//
//	// ImportCollection 允许未跟踪的 collection，其余必须已跟踪
//	if event.EventType != ImportCollection {
//		if _, tracked := om.collectionOrders[strings.ToLower(event.CollectionAddr)]; !tracked {
//			om.logger.WarnContext(ctx, "untracked collection",
//				"collection_addr", event.CollectionAddr,
//			)
//			return TradeEvent{}, false
//		}
//	}
//
//	// 非阻塞通知，防止 channel 满时阻塞消费循环
//	select {
//	case om.collectionListedCh <- event.CollectionAddr:
//	default:
//		om.logger.WarnContext(ctx, "collectionListedCh is full, dropping notification",
//			"collection_addr", event.CollectionAddr,
//		)
//	}
//
//	return event, true
//}
//
//func (om *OrderManager) clearRemainingEvents(ctx context.Context, key string) {
//	if err := om.Xkv.Redis.Ltrim(ctx, key, 1, 0).Err(); err != nil {
//		om.logger.ErrorContext(ctx, "failed on flush remaining trade events", "error", err)
//	}
//}
//
//// handleListing 上架事件
//func (om *OrderManager) handleListing(ctx context.Context, _ *OrderManager, event TradeEvent) error {
//	tradeInfo := om.collectionOrders[strings.ToLower(event.CollectionAddr)]
//	_, maxPrice := tradeInfo.orders.GetMax()
//
//	if tradeInfo.orders.Len() == 0 || maxPrice.GreaterThan(event.Price) {
//		tradeInfo.orders.Add(event.OrderId, event.Price, event.From, event.TokenID)
//	}
//
//	return om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr)
//}
//
//// handleCancelOrExpired 取消/过期事件
//func (om *OrderManager) handleCancelOrExpired(ctx context.Context, _ *OrderManager, event TradeEvent) error {
//	tradeInfo := om.collectionOrders[strings.ToLower(event.CollectionAddr)]
//	tradeInfo.orders.Remove(event.OrderId)
//
//	if tradeInfo.orders.Len() == 0 {
//		if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
//			return fmt.Errorf("reload after cancel/expired: %w", err)
//		}
//	}
//
//	return om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr)
//}
//
//// handleBuyOrTransfer 购买/转移事件
//func (om *OrderManager) handleBuyOrTransfer(ctx context.Context, _ *OrderManager, event TradeEvent) error {
//	tradeInfo := om.collectionOrders[strings.ToLower(event.CollectionAddr)]
//
//	if event.EventType == Buy {
//		tradeInfo.orders.Remove(event.OrderId)
//	}
//
//	if tradeInfo.orders.Len() == 0 {
//		if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
//			return fmt.Errorf("reload after buy/transfer (empty queue): %w", err)
//		}
//		return om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr)
//	}
//
//	tradeInfo.orders.RemoveMakerOrders(event.From, event.TokenID)
//
//	validOrders, err := om.getUserValidOrders(ctx, event.CollectionAddr, event.TokenID, event.To)
//	if err != nil {
//		return fmt.Errorf("get user valid orders: %w", err)
//	}
//
//	for _, order := range validOrders {
//		if order.IsOpenseaBanned {
//			continue
//		}
//		_, maxPrice := tradeInfo.orders.GetMax()
//		if maxPrice.GreaterThan(order.Price) {
//			tradeInfo.orders.Add(order.OrderID, order.Price, order.Maker, order.TokenId)
//		}
//	}
//
//	if tradeInfo.orders.Len() == 0 {
//		if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
//			return fmt.Errorf("reload after adding buyer orders: %w", err)
//		}
//	}
//
//	return om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr)
//}
//
//// handleImportCollection 导入新 Collection
//func (om *OrderManager) handleImportCollection(ctx context.Context, _ *OrderManager, event TradeEvent) error {
//	addrLower := strings.ToLower(event.CollectionAddr)
//	if _, exists := om.collectionOrders[addrLower]; exists {
//		om.logger.WarnContext(ctx, "import collection repeated",
//			"collection_addr", event.CollectionAddr,
//		)
//		return nil
//	}
//
//	om.collectionOrders[addrLower] = &collectionTradeInfo{
//		floorPrice: decimal.Zero,
//		orders:     NewPriorityQueueMap(maxQueueLength),
//	}
//	return nil
//}
//
//// handleUpdateCollection 更新 Collection
//func (om *OrderManager) handleUpdateCollection(ctx context.Context, _ *OrderManager, event TradeEvent) error {
//	tradeInfo := om.collectionOrders[strings.ToLower(event.CollectionAddr)]
//	_, currentFloor := tradeInfo.orders.GetMin()
//
//	if currentFloor.Equal(event.Price) {
//		return nil
//	}
//
//	if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
//		return fmt.Errorf("reload on update collection: %w", err)
//	}
//
//	return om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr)
//}
