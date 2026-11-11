package ordermanager

import (
	"MiSwapSync/stores/gdb"
	"MiSwapSync/stores/gdb/multi"
	"MiSwapSync/stores/xkv"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"log/slog"
	"strings"
	"time"
)

const (
	MaxBatchReqNum = 100
	maxQueueLength = 100
)

const CacheTradeEventsQueuePre = "cache:ms:trade:events:%s"

type ValidOrder struct {
	multi.Order
	IsOpenseaBanned bool `json:"is_opensea_banned"`
}

func (om *OrderManager) floorPriceProcess(ctx context.Context) {
	key := genTradeEventsCacheKey(om.chain)
	if err := om.Xkv.Redis.Ltrim(key, 1, 0); err != nil {
		slog.WarnContext(ctx, "failed to flush remaining trade events", "error", err)
	}
	// 从数据库加载订单并更新地板价
	if err := om.loadCollectionTradeInfo(ctx); err != nil {
		slog.ErrorContext(ctx, "[floorPrice] load orders to queue", "error", err)
		return
	}
	// 持续监听并处理交易事件
	for {
		select {
		case <-ctx.Done():
			slog.InfoContext(ctx, "floorPriceProcess cancelled", "error", ctx.Err())
			return
		default:
		}
		// 从缓存中获取交易事件
		result, err := om.Xkv.Redis.Lpop(key)
		if err != nil || result == "" {
			if err != nil && !errors.Is(err, redis.Nil) {
				slog.WarnContext(ctx, "failed to get trade events from cache")
			}
			// 等待或取消
			select {
			case <-ctx.Done():
				return
			case <-time.After(1 * time.Second):
			}
			continue
		}
		slog.InfoContext(ctx, "get trade events from cache", "event_content", result)
		// 解析事件内容
		var event TradeEvent
		if err := json.Unmarshal([]byte(result), &event); err != nil {
			slog.WarnContext(ctx, "failed to unmarshal trade event info", "error", err.Error())
			continue
		}

		// 检查 collection 是否被跟踪
		tradeInfo, ok := om.collectionOrders[strings.ToLower(event.CollectionAddr)]
		if !ok && event.EventType != ImportCollection {
			slog.WarnContext(ctx, "untracked collection", "collection_addr", event.CollectionAddr)
			continue
		}

		// 通知 collection 状态更新
		if event.CollectionAddr != "" {
			select {
			case om.collectionListedCh <- event.CollectionAddr:
			default:
			}
		}
		// 根据不同事件类型分发处理
		switch event.EventType {
		case Listing:
			om.handleListing(ctx, event, tradeInfo)
		case Cancel, Expired:
			om.handleCancelOrExpired(ctx, event, tradeInfo)
		case Buy, Transfer:
			om.handleBuyOrTransfer(ctx, event, tradeInfo)
		case ImportCollection:
			om.handleImportCollection(ctx, event)
		case UpdateCollection:
			om.handleUpdateCollection(ctx, event, tradeInfo)
		default:
			slog.ErrorContext(ctx, "unsupported event type", "event_content", result)
		}
	}
}

// logError 统一错误日志记录（减少重复代码）
func (om *OrderManager) logError(ctx context.Context, event *TradeEvent, err error, extra ...any) {
	fields := []any{
		"event_type", event.EventType,
		"order_id", event.OrderId,
		"collection_addr", event.CollectionAddr,
		"error", err.Error(),
	}
	fields = append(fields, extra...)
	slog.ErrorContext(ctx, "trade event handler error", fields...)
}

// handleListing 处理上架事件
func (om *OrderManager) handleListing(ctx context.Context, event TradeEvent, tradeInfo *collectionTradeInfo) {
	// 只有当价格低于队列最高价或队列为空时才添加订单
	_, price := tradeInfo.orders.GetMax()
	if price.GreaterThan(event.Price) || tradeInfo.orders.Len() == 0 {
		tradeInfo.orders.Add(event.OrderId, event.Price, event.From, event.TokenID)
	}
	// 检查并更新地板价
	if err := om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr); err != nil {
		slog.ErrorContext(ctx, "failed to update floor_price", "price", event.Price.String())
		om.logError(ctx, &event, err, "floor_price", event.Price.String())
	}
}

// checkAndUpdateFloorPrice 检查并更新地板价（如果发生变化）
func (om *OrderManager) checkAndUpdateFloorPrice(ctx context.Context, address string) error {
	tradeInfo, ok := om.collectionOrders[strings.ToLower(address)]
	if !ok {
		slog.WarnContext(ctx, "untracked collection", "collection_addr", address)
		return fmt.Errorf("untracked collection: %s", address)
	}

	_, newFloorPrice := tradeInfo.orders.GetMin()

	if !newFloorPrice.Equal(tradeInfo.floorPrice) {
		tradeInfo.floorPrice = newFloorPrice
		om.collectionOrders[strings.ToLower(address)] = tradeInfo

		if err := om.updateFloorPrice(ctx, address, newFloorPrice); err != nil {
			return fmt.Errorf("failed to update collection floor price: %w", err)
		}

		slog.InfoContext(ctx, "update collection floor price",
			"collection_addr", address, "floor_price", newFloorPrice.String())
	}
	return nil
}

// handleCancelOrExpired 处理取消或过期事件
func (om *OrderManager) handleCancelOrExpired(ctx context.Context, event TradeEvent, tradeInfo *collectionTradeInfo) {
	tradeInfo.orders.Remove(event.OrderId)
	if tradeInfo.orders.Len() == 0 {
		if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
			om.logError(ctx, &event, err)
			return
		}
	}
	if err := om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr); err != nil {
		om.logError(ctx, &event, err)
	}
}

// handleBuyOrTransfer 处理购买或转移事件
func (om *OrderManager) handleBuyOrTransfer(ctx context.Context, event TradeEvent, tradeInfo *collectionTradeInfo) {
	// 如果是购买事件，从队列中删除订单
	if event.EventType == Buy {
		tradeInfo.orders.Remove(event.OrderId)
	}

	// 检查队列是否为空，为空则重新加载订单
	if tradeInfo.orders.Len() == 0 {
		if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
			om.logError(ctx, &event, err)
			return
		}
	} else {
		// 移除卖家的所有订单
		tradeInfo.orders.RemoveMakerOrders(event.From, event.TokenID)

		// 获取买家的有效订单
		orders, err := om.getUserValidOrders(ctx, event.CollectionAddr, event.TokenID, event.To)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get users valid orders",
				"event_type", event.EventType,
				"order_id", event.OrderId,
				"collection_addr", event.CollectionAddr,
				"from", event.From,
				"to", event.To,
				"error", err.Error())
			return
		}

		// 添加买家的有效订单到队列
		for _, order := range orders {
			if !order.IsOpenseaBanned {
				_, maxPrice := tradeInfo.orders.GetMax()
				if maxPrice.GreaterThan(order.Price) {
					tradeInfo.orders.Add(order.OrderID, order.Price, order.Maker, order.TokenId)
				}
			}
		}

		// 如果队列为空，重新加载订单
		if tradeInfo.orders.Len() == 0 {
			if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
				om.logError(ctx, &event, err)
				return
			}
		}
	}
	if err := om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr); err != nil {
		om.logError(ctx, &event, err, "floor_price", event.Price.String())
	}
}

// handleImportCollection 处理导入新 Collection 事件
func (om *OrderManager) handleImportCollection(ctx context.Context, event TradeEvent) {
	if _, ok := om.collectionOrders[strings.ToLower(event.CollectionAddr)]; ok {
		slog.WarnContext(ctx, "import collection repeated", "collection_addr", event.CollectionAddr)
		return
	}

	om.collectionOrders[strings.ToLower(event.CollectionAddr)] = &collectionTradeInfo{
		floorPrice: decimal.Zero,
		orders:     NewPriorityQueueMap(maxQueueLength),
	}
}

// handleUpdateCollection 处理更新 Collection 事件
func (om *OrderManager) handleUpdateCollection(ctx context.Context, event TradeEvent, tradeInfo *collectionTradeInfo) {
	// 检查地板价是否变化
	_, floorPrice := tradeInfo.orders.GetMin()
	if floorPrice.Equal(event.Price) {
		return
	}

	if err := om.reloadCollectionOrders(ctx, event.CollectionAddr); err != nil {
		om.logError(ctx, &event, err)
		return
	}

	if err := om.checkAndUpdateFloorPrice(ctx, event.CollectionAddr); err != nil {
		om.logError(ctx, &event, err, "floor_price", event.Price.String())
	}
}

// loadCollectionTradeInfo 初始化并加载所有集合的交易信息
func (om *OrderManager) loadCollectionTradeInfo(ctx context.Context) error {
	// 1. 从数据库加载所有集合信息
	var collections []*multi.Collection
	if err := om.DB.WithContext(ctx).Table(gdb.GetMultiProjectCollectionTableName(om.project, om.chain)).
		Select("id, address, floor_price").Where("1=1").Scan(&collections).Error; err != nil {
		return fmt.Errorf("failed to get collection floor price: %w", err)
	}

	var collectionAddrs []string
	for _, collection := range collections {
		collectionAddrs = append(collectionAddrs, collection.Address)
		om.collectionOrders[strings.ToLower(collection.Address)] = &collectionTradeInfo{
			floorPrice: collection.FloorPrice,
			orders:     NewPriorityQueueMap(maxQueueLength),
		}
	}

	// 2. 分批加载每个集合的有效订单
	var totalOrders []*multi.Order
	var id int64
	for {
		var orders []*multi.Order
		if err := om.DB.WithContext(ctx).Table(fmt.Sprintf("%s as co", gdb.GetMultiProjectOrderTableName(om.project, om.chain))).
			Select("co.id as id, co.order_id as order_id, co.collection_address as collection_address, co.price as price, co.maker as maker, co.token_id as token_id").
			Joins(fmt.Sprintf("join %s ci on co.collection_address = ci.collection_address and co.token_id = ci.token_id", gdb.GetMultiProjectItemTableName(om.project, om.chain))).
			Where("co.order_type=? and co.order_status = ? and co.maker = ci.owner and (ci.is_opensea_banned, co.marketplace_id) != (true, 1) and co.id > ?",
				multi.ListingType, multi.OrderStatusActive, id).
			Order("co.id asc").Limit(1000).
			Scan(&orders).Error; err != nil {
			return fmt.Errorf("failed to get collection orders: %w", err)
		}
		totalOrders = append(totalOrders, orders...)
		if len(orders) < 1000 {
			break
		}
		id = orders[len(orders)-1].ID
	}

	// 3. 将订单添加到对应集合的优先队列中
	for _, order := range totalOrders {
		ordersQueue, ok := om.collectionOrders[strings.ToLower(order.CollectionAddress)]
		if !ok {
			slog.WarnContext(ctx, "untracked collection", "collection_addr", order.CollectionAddress)
			continue
		}
		ordersQueue.orders.Add(order.OrderID, order.Price, order.Maker, order.TokenId)
	}

	// 4. 检查并更新每个集合的地板价
	for addr, tradeInfo := range om.collectionOrders {
		_, floorPrice := tradeInfo.orders.GetMin()
		if !floorPrice.Equal(tradeInfo.floorPrice) {
			om.collectionOrders[strings.ToLower(addr)].floorPrice = floorPrice
			if err := om.updateFloorPrice(ctx, addr, floorPrice); err != nil {
				slog.WarnContext(ctx, "failed to update collection floor price",
					"collection_addr", addr, "error", err.Error())
			}
		}
	}
	return nil
}

// updateFloorPrice 更新数据库中的地板价
func (om *OrderManager) updateFloorPrice(ctx context.Context, collectionAddr string, price decimal.Decimal) error {
	if err := om.DB.WithContext(ctx).Table(gdb.GetMultiProjectCollectionTableName(om.project, om.chain)).
		Where("address=?", collectionAddr).Update("floor_price", price).Error; err != nil {
		return fmt.Errorf("failed to update collection floor price: %w", err)
	}
	return nil
}

func (om *OrderManager) getUserValidOrders(ctx context.Context, address, tokenID, maker string) ([]*ValidOrder, error) {
	var orders []*ValidOrder
	if err := om.DB.WithContext(ctx).Table(fmt.Sprintf("%s as co", gdb.GetMultiProjectOrderTableName(om.project, om.chain))).
		Select("co.id as id, co.order_id as order_id, co.maker as maker, ci.is_opensea_banned as is_opensea_banned, co.collection_address as collection_address, co.price as price, co.token_id as token_id").
		Joins(fmt.Sprintf("join %s ci on co.collection_address = ci.collection_address and co.token_id = ci.token_id", gdb.GetMultiProjectItemTableName(om.project, om.chain))).
		Where("co.order_type=? and co.order_status = ? and co.maker = ci.owner and (ci.is_opensea_banned, co.marketplace_id) != (true, 1)",
			multi.ListingType, multi.OrderStatusActive).
		Where("co.collection_address = ? and co.token_id = ? and co.maker = ?", address, tokenID, maker).
		Order("co.price asc").Limit(100).
		Scan(&orders).Error; err != nil {
		return nil, fmt.Errorf("failed to get user orders: %w", err)
	}
	return orders, nil
}

// getLowestPrice100Orders 获取指定集合价格最低的 100 个订单
func (om *OrderManager) getLowestPrice100Orders(ctx context.Context, address string) ([]*multi.Order, error) {
	var orders []*multi.Order
	if err := om.DB.WithContext(ctx).Table(fmt.Sprintf("%s as co", gdb.GetMultiProjectOrderTableName(om.project, om.chain))).
		Select("co.id, co.order_id as order_id, co.collection_address, co.price, co.maker, co.token_id").
		Joins(fmt.Sprintf("join %s ci on co.collection_address = ci.collection_address and co.token_id = ci.token_id", gdb.GetMultiProjectItemTableName(om.project, om.chain))).
		Where("co.order_type=? and co.order_status = ? and co.maker = ci.owner and (ci.is_opensea_banned, co.marketplace_id) != (true, 1)",
			multi.ListingType, multi.OrderStatusActive).
		Where("co.collection_address = ?", address).
		Order("co.price asc").Limit(100).
		Scan(&orders).Error; err != nil {
		return nil, fmt.Errorf("failed to get collection lowest price orders: %w", err)
	}
	return orders, nil
}

// reloadCollectionOrders 重新加载指定集合的订单信息
func (om *OrderManager) reloadCollectionOrders(ctx context.Context, address string) error {
	tradeInfo, ok := om.collectionOrders[strings.ToLower(address)]
	if !ok {
		slog.WarnContext(ctx, "untracked collection", "collection_addr", address)
		return fmt.Errorf("untracked collection: %s", address)
	}

	orders, err := om.getLowestPrice100Orders(ctx, address)
	if err != nil {
		return fmt.Errorf("failed to get 100 orders at the lowest price: %w", err)
	}

	// 重新构建优先级队列
	tradeInfo.orders = NewPriorityQueueMap(maxQueueLength)
	for _, order := range orders {
		tradeInfo.orders.Add(order.OrderID, order.Price, order.Maker, order.TokenId)
	}
	om.collectionOrders[strings.ToLower(address)] = tradeInfo
	return nil
}

// genTradeEventsCacheKey 生成交易事件缓存 key
func genTradeEventsCacheKey(chain string) string {
	return fmt.Sprintf(CacheTradeEventsQueuePre, chain)
}

func AddUpdatePriceEvent(kv *xkv.Store, event *TradeEvent, chain string) error {
	// 验证事件合法性
	// 如果是非Transfer/ImportCollection/UpdateCollection事件,必须有OrderId
	if event.EventType != Transfer && event.EventType != ImportCollection && event.EventType != UpdateCollection && event.OrderId == "" {
		return errors.New("invalid update collection floor price. event order id is null")
	}
	// 如果是Listing事件,价格不能为0
	if event.EventType == Listing && event.Price.IsZero() {
		return errors.New("invalid update collection floor price. price is 0")
	}
	// 如果是Listing事件,必须有TokenID
	if event.EventType == Listing && event.TokenID == "" {
		return errors.New("invalid update collection floor price. token_id is null")
	}

	// 序列化事件
	rawEvent, err := json.Marshal(event)
	if err != nil {
		return errors.Wrap(err, "failed to marshal event")
	}

	// 获取Redis队列key
	key := genTradeEventsCacheKey(chain)
	// 将事件添加到Redis队列
	if _, err := kv.Rpush(key, string(rawEvent)); err != nil {
		return errors.Wrap(err, "failed to push trade event to queue")
	}
	return nil
}
