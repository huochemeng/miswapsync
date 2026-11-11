package orderbookindex

import (
	"MiSwapSync/chain/chainclient"
	miswap "MiSwapSync/chain/nftchainservice/bind_erc721"
	"MiSwapSync/chain/types"
	"MiSwapSync/config"
	"MiSwapSync/ordermanager"
	"MiSwapSync/stores/gdb"
	"MiSwapSync/stores/gdb/base"
	"MiSwapSync/stores/gdb/multi"
	"MiSwapSync/stores/xkv"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethereumTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/threading"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log/slog"
	"math/big"
	"strings"
	"time"
)

const (
	EventIndexType   = 6
	SleepInterval    = 10 // in seconds
	SyncBlockPeriod  = 10
	LogMakeTopic     = "0xfc37f2ff950f95913eb7182357ba3c14df60ef354bc7d6ab1ba2815f249fffe6"
	LogCancelTopic   = "0x0ac8bb53fac566d7afc05d8b4df11d7690a7b27bdc40b54e4060f9b21fb849bd"
	LogMatchTopic    = "0xf629aecab94607bc43ce4aebd564bf6e61c7327226a797b002de724b9944b20e"
	FixForCollection = 0
	FixForItem       = 1
	List             = 0
	Bid              = 1
	HexPrefix        = "0x"
	ZeroAddress      = "0x0000000000000000000000000000000000000000"
)

const (
	DBBatchSizeLimit                 = 200
	CollectionFloorChangeIndexType   = 5
	CollectionFloorSyncPeriod        = 150                // in seconds
	DaySeconds                       = 3600 * 24          // in seconds
	MaxCollectionFloorTimeDifference = 10                 // in seconds
	CollectionFloorTimeRange         = 3600 * 24 * 30 * 2 // in seconds
)

type Order struct {
	Side     uint8
	SaleKind uint8
	Maker    common.Address
	Nft      struct {
		TokenId        *big.Int
		CollectionAddr common.Address
		Amount         *big.Int
	}
	Price  *big.Int
	Expiry uint64
	Salt   uint64
}

type Service struct {
	db           *gorm.DB
	kv           *xkv.Store
	orderManager *ordermanager.OrderManager
	chainClient  chainclient.ChainClient
	chainId      int64
	chain        string
	parsedAbi    abi.ABI
	dexAddress   string
	ethAddress   string
	projectName  string
}

var MultiChainMaxBlockDifference = map[string]uint64{
	"eth":        1,
	"optimism":   2,
	"starknet":   1,
	"arbitrum":   2,
	"base":       2,
	"zksync-era": 2,
}

func New(ctx context.Context, cfg *config.Config, db *gorm.DB, xkv *xkv.Store, chainClient chainclient.ChainClient, chainId int64, chain string, orderManager *ordermanager.OrderManager) *Service {
	parsedAbi, _ := abi.JSON(strings.NewReader(miswap.MiSwapOrderBookMetaData.ABI)) // 通过ABI实例化
	return &Service{
		db:           db,
		kv:           xkv,
		chainClient:  chainClient,
		orderManager: orderManager,
		chain:        chain,
		chainId:      chainId,
		parsedAbi:    parsedAbi,
		dexAddress:   cfg.ContractCfg.DexAddress,
		ethAddress:   cfg.ContractCfg.EthAddress,
		projectName:  cfg.ProjectCfg.Name,
	}
}

func (s *Service) Start(ctx context.Context) {
	// 同步数据，数据存入数据库
	threading.GoSafe(func() { s.SyncOrderBookEventLoop(ctx) })
	// 计算地板价的服务
	threading.GoSafe(func() { s.UpKeepingCollectionFloorChangeLoop(ctx) })
}

func (s *Service) SyncOrderBookEventLoop(ctx context.Context) {
	var indexedStatus base.IndexedStatus
	if err := s.db.WithContext(ctx).Table(base.IndexedStatusTableName()).
		Where("chain_id = ? and index_type = ?", s.chainId, EventIndexType).
		First(&indexedStatus).Error; err != nil {
		slog.ErrorContext(ctx, "failed to get listing index status", "error", err)
		return
	}
	// 查找 indexed_status.LastIndexedBlock
	lastSyncBlock := uint64(indexedStatus.LastIndexedBlock)
	//使用一个ticker，用来计时
	ticker := time.NewTicker(SleepInterval * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			slog.WarnContext(ctx, "SyncOrderBookEventLoop stopped due to context cancellation")
			return
		default:
		}
		currentBlockNum, err := s.chainClient.BlockNumber() // 以轮询的方式获取当前区块高度
		if err != nil {
			slog.ErrorContext(ctx, "failed to get current block number", "error", err)
			select {
			case <-ctx.Done():
				slog.WarnContext(ctx, "SyncOrderBookEventLoop stopped due to context cancellation")
				return
			case <-ticker.C:
				continue
			}
		}

		// 如果上次同步的区块高度大于当前区块高度，等待一段时间后再次轮询
		if lastSyncBlock > currentBlockNum-MultiChainMaxBlockDifference[s.chain] {
			time.Sleep(SleepInterval * time.Second)
			continue
		}

		startBlock := lastSyncBlock
		endBlock := startBlock + SyncBlockPeriod
		if endBlock > currentBlockNum-MultiChainMaxBlockDifference[s.chain] { // 如果结束区块高度大于当前区块高度，将结束区块高度设置为当前区块高度
			endBlock = currentBlockNum - MultiChainMaxBlockDifference[s.chain]
		}

		query := types.FilterQuery{
			FromBlock: new(big.Int).SetUint64(startBlock),
			ToBlock:   new(big.Int).SetUint64(endBlock),
			Addresses: []string{s.dexAddress},
		}
		// 真正获取 日志的方法
		logs, err := s.chainClient.FilterLogs(ctx, query) //同时获取多个（SyncBlockPeriod）区块的日志
		if err != nil {
			slog.WarnContext(ctx, "failed to get log", "error", err)
			time.Sleep(SleepInterval * time.Second)
			continue
		}

		for _, log := range logs { // 遍历日志，根据不同的topic处理不同的事件
			ethLog := log.(ethereumTypes.Log)
			switch ethLog.Topics[0].String() {
			case LogMakeTopic:
				s.handleMakeEvent(ctx, ethLog)
				// 订单生成
			case LogCancelTopic:
				s.handleCancelEvent(ctx, ethLog)
			// 订单取消
			case LogMatchTopic:
				s.handleMatchEvent(ctx, ethLog)
			// 订单成交
			default:
			}
		}

		lastSyncBlock = endBlock + 1 // 更新最后同步的区块高度
		if err := s.db.WithContext(ctx).Table(base.IndexedStatusTableName()).
			Where("chain_id = ? and index_type = ?", s.chainId, EventIndexType).
			Update("last_indexed_block", lastSyncBlock).Error; err != nil {
			slog.WarnContext(ctx, "failed to update order book event sync block number", "error", err)
			return
		}

		slog.InfoContext(ctx, "sync order book event ...", "start_block", startBlock, "end_block", endBlock)
	}
}

func (s *Service) handleMakeEvent(ctx context.Context, log ethereumTypes.Log) {
	var event struct {
		OrderKey [32]byte
		Nft      struct {
			TokenId        *big.Int
			CollectionAddr common.Address
			Amount         *big.Int
		}
		Price  *big.Int
		Expiry uint64
		Salt   uint64
	}

	// 通过ABI解析日志数据
	err := s.parsedAbi.UnpackIntoInterface(&event, "LogMake", log.Data)
	if err != nil {
		slog.ErrorContext(ctx, "failed to unpack LogMake event:", "error", err)
		return
	}
	// Extract indexed fields from topics
	side := uint8(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())
	saleKind := uint8(new(big.Int).SetBytes(log.Topics[2].Bytes()).Uint64())
	maker := common.BytesToAddress(log.Topics[3].Bytes())

	var orderType int64
	if side == Bid { // 买单
		if saleKind == FixForCollection { // 针对集合的买单
			orderType = multi.CollectionBidOrder
		} else { // 针对某个具体NFT的买单
			orderType = multi.ItemBidOrder
		}
	} else { // 卖单
		orderType = multi.ListingOrder
	}
	newOrder := multi.Order{
		CollectionAddress: event.Nft.CollectionAddr.String(),
		MarketplaceId:     multi.MarketOrderBook,
		TokenId:           event.Nft.TokenId.String(),
		OrderID:           HexPrefix + hex.EncodeToString(event.OrderKey[:]),
		OrderStatus:       multi.OrderStatusActive,
		EventTime:         time.Now().Unix(),
		ExpireTime:        int64(event.Expiry),
		CurrencyAddress:   s.ethAddress,
		Price:             decimal.NewFromBigInt(event.Price, 0),
		Maker:             maker.String(),
		Taker:             ZeroAddress,
		QuantityRemaining: event.Nft.Amount.Int64(),
		Size:              event.Nft.Amount.Int64(),
		OrderType:         orderType,
		Salt:              int64(event.Salt),
	}
	if err := s.db.WithContext(ctx).Table(multi.OrderTableName(s.chain)).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&newOrder).Error; err != nil {
		slog.ErrorContext(ctx, "failed to create order", "error", err)
	}
	blockTime, err := s.chainClient.BlockTimeByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
	if err != nil {
		slog.ErrorContext(ctx, "failed to get block time", "error", err)
		return
	}
	var activityType int
	if side == Bid {
		if saleKind == FixForCollection {
			activityType = multi.CollectionBid
		} else {
			activityType = multi.ItemBid
		}
	} else {
		activityType = multi.Listing
	}
	newActivity := multi.Activity{ // 将订单信息存入活动表
		ActivityType:      activityType,
		Maker:             maker.String(),
		Taker:             ZeroAddress,
		MarketplaceID:     multi.MarketOrderBook,
		CollectionAddress: event.Nft.CollectionAddr.String(),
		TokenId:           event.Nft.TokenId.String(),
		CurrencyAddress:   s.ethAddress,
		Price:             decimal.NewFromBigInt(event.Price, 0),
		BlockNumber:       int64(log.BlockNumber),
		TxHash:            log.TxHash.String(),
		EventTime:         int64(blockTime),
	}
	if err := s.db.WithContext(ctx).Table(multi.ActivityTableName(s.chain)).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&newActivity).Error; err != nil {
		slog.ErrorContext(ctx, "failed to crate activity", "error", err)
	}
	if err := s.orderManager.AddToOrderManagerQueue(&multi.Order{ // 将订单信息存入订单管理队列
		ExpireTime:        newOrder.ExpireTime,
		OrderID:           newOrder.OrderID,
		CollectionAddress: newOrder.CollectionAddress,
		TokenId:           newOrder.TokenId,
		Price:             newOrder.Price,
		Maker:             newOrder.Maker,
	}); err != nil {
		slog.ErrorContext(ctx, "failed to add order to manager queue", "error", err, "order_id", newOrder.OrderID)
	}
}

// logMatchEventData 提取事件数据结构，方便在子函数间传递
type logMatchEventData struct {
	MakeOrder Order
	TakeOrder Order
	FillPrice *big.Int
}

// handleMatchEvent 入口函数：只负责解析、编排和顶层错误处理
func (s *Service) handleMatchEvent(ctx context.Context, log ethereumTypes.Log) {
	// 1. 解析事件 Data
	var event logMatchEventData
	if err := s.parsedAbi.UnpackIntoInterface(&event, "LogMatch", log.Data); err != nil {
		slog.ErrorContext(ctx, "failed to unpacking LogMatch event,", "error", err)
		return
	}
	// 2. 解析 Topics 中的 OrderId
	makeOrderId := HexPrefix + hex.EncodeToString(log.Topics[1].Bytes())
	takeOrderId := HexPrefix + hex.EncodeToString(log.Topics[2].Bytes())

	// 3. 根据买卖方向解析角色映射
	sellOrderId, buyOrderId, owner, collection, tokenId, from, to := s.resolveMatchRoles(&event, makeOrderId, takeOrderId)

	// 4. 更新撮合双方的订单状态
	if err := s.updateMatchedOrders(ctx, sellOrderId, buyOrderId, to); err != nil {
		return
	}

	// 5. 记录交易活动 (Activity)
	s.recordMatchActivity(ctx, &event, &log, collection, tokenId)

	// 6. 更新 NFT 所有者
	if err := s.updateItemOwner(ctx, collection, tokenId, owner); err != nil {
		slog.ErrorContext(ctx, "failed to update item owner")
		return
	}

	// 7. 推送价格更新事件
	s.pushPriceUpdateEvent(sellOrderId, collection, tokenId, from, to)
}

// resolveMatchRoles 纯函数：根据 Side 决定买卖角色
func (s *Service) resolveMatchRoles(event *logMatchEventData, makeOrderId, takeOrderId string) (
	sellOrderId, buyOrderId, owner, collection, tokenId, from, to string,
) {
	if event.MakeOrder.Side == Bid {
		// 买单，由卖方发起交易撮合
		return takeOrderId, makeOrderId,
			strings.ToLower(event.MakeOrder.Maker.String()),
			event.TakeOrder.Nft.CollectionAddr.String(),
			event.TakeOrder.Nft.TokenId.String(),
			event.TakeOrder.Maker.String(),
			event.MakeOrder.Maker.String()
	}

	// 卖单，由买方发起交易撮合
	return makeOrderId, takeOrderId,
		strings.ToLower(event.TakeOrder.Maker.String()),
		event.MakeOrder.Nft.CollectionAddr.String(),
		event.MakeOrder.Nft.TokenId.String(),
		event.MakeOrder.Maker.String(),
		event.TakeOrder.Maker.String()
}

// updateMatchedOrders 封装对称的订单状态更新逻辑
func (s *Service) updateMatchedOrders(ctx context.Context, sellOrderId, buyOrderId, taker string) error {
	table := multi.OrderTableName(s.chain)
	// 1. 更新卖方订单状态为已成交
	if err := s.db.WithContext(ctx).Table(table).
		Where("order_id = ?", sellOrderId).
		Updates(map[string]interface{}{
			"order_status":       multi.OrderStatusFilled,
			"quantity_remaining": 0,
			"taker":              taker,
		}).Error; err != nil {
		slog.ErrorContext(ctx, "failed to update sell order status", "order_id", sellOrderId)
		return err
	}

	// 2. 查询买方订单信息
	var buyOrder multi.Order
	if err := s.db.WithContext(ctx).Table(table).
		Where("order_id = ?", buyOrderId).
		First(&buyOrder).Error; err != nil {
		slog.ErrorContext(ctx, "failed to get buy order")
		return err
	}

	// 3. 更新买方订单的剩余数量或标记完成
	if buyOrder.QuantityRemaining > 1 {
		if err := s.db.WithContext(ctx).Table(table).
			Where("order_id = ?", buyOrderId).
			Update("quantity_remaining", buyOrder.QuantityRemaining-1).Error; err != nil {
			slog.ErrorContext(ctx, "failed to update buy order quantity_remaining")
			return err
		}
	} else {
		if err := s.db.WithContext(ctx).Table(table).
			Where("order_id = ?", buyOrderId).
			Updates(map[string]interface{}{
				"order_status":       multi.OrderStatusFilled,
				"quantity_remaining": 0,
			}).Error; err != nil {
			slog.ErrorContext(ctx, "failed to update buy order status")
			return err
		}
	}
	return nil
}

// recordMatchActivity 负责获取区块时间并写入 Activity 表
func (s *Service) recordMatchActivity(ctx context.Context, event *logMatchEventData, log *ethereumTypes.Log, collection, tokenId string) {
	blockTime, err := s.chainClient.BlockTimeByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
	if err != nil {
		slog.ErrorContext(ctx, "failed to get block time")
		return
	}

	newActivity := multi.Activity{
		ActivityType:      multi.Sale,
		Maker:             event.MakeOrder.Maker.String(),
		Taker:             event.TakeOrder.Maker.String(),
		MarketplaceID:     multi.MarketOrderBook,
		CollectionAddress: collection,
		TokenId:           tokenId,
		CurrencyAddress:   s.ethAddress,
		Price:             decimal.NewFromBigInt(event.FillPrice, 0),
		BlockNumber:       int64(log.BlockNumber),
		TxHash:            log.TxHash.String(),
		EventTime:         int64(blockTime),
	}

	if err := s.db.WithContext(ctx).Table(multi.ActivityTableName(s.chain)).Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&newActivity).Error; err != nil {
		slog.WarnContext(ctx, "failed to create activity")
	}
}

// updateItemOwner 负责更新 NFT 归属
func (s *Service) updateItemOwner(ctx context.Context, collection, tokenId, owner string) error {
	return s.db.WithContext(ctx).Table(multi.ItemTableName(s.chain)).
		Where("collection_address = ? and token_id = ?", strings.ToLower(collection), tokenId).
		Update("owner", owner).Error
}

// pushPriceUpdateEvent 负责向消息队列推送价格更新事件
func (s *Service) pushPriceUpdateEvent(sellOrderId, collection, tokenId, from, to string) {
	tradeEvent := &ordermanager.TradeEvent{
		OrderId:        sellOrderId,
		CollectionAddr: collection,
		EventType:      ordermanager.Buy,
		TokenID:        tokenId,
		From:           from,
		To:             to,
	}

	if err := ordermanager.AddUpdatePriceEvent(s.kv, tradeEvent, s.chain); err != nil {
		slog.Error("failed to add update price event", "error", err, "type", "sale", "order_id", sellOrderId)
	}
}

// handleCancelEvent 入口函数：仅负责编排取消事件的处理流程
func (s *Service) handleCancelEvent(ctx context.Context, log ethereumTypes.Log) {
	// 1. 从 Topics 中提取 orderId
	orderId := HexPrefix + hex.EncodeToString(log.Topics[1].Bytes())

	// 2. 更新订单状态为已取消
	if err := s.cancelOrder(ctx, orderId); err != nil {
		slog.ErrorContext(ctx, "failed to update order status")
		return
	}
	// 3. 查询被取消的订单详情
	cancelOrder, err := s.getOrderByID(ctx, orderId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get cancel order")
		return
	}

	// 4. 构建并写入取消活动记录
	if err := s.saveCancelActivity(ctx, cancelOrder, &log); err != nil {
		slog.ErrorContext(ctx, "failed to create activity")
	}

	// 5. 推送价格更新事件
	s.pushCancelPriceEvent(cancelOrder)
}

// ---------- 数据访问层 ----------
func (s *Service) cancelOrder(ctx context.Context, orderId string) error {
	return s.db.WithContext(ctx).
		Table(multi.OrderTableName(s.chain)).
		Where("order_id = ?", orderId).
		Update("order_status", multi.OrderStatusCancelled).Error
}

func (s *Service) getOrderByID(ctx context.Context, orderId string) (multi.Order, error) {
	var order multi.Order
	err := s.db.WithContext(ctx).
		Table(multi.OrderTableName(s.chain)).
		Where("order_id = ?", orderId).
		First(&order).Error
	return order, err
}

// ---------- 类型映射层（纯函数） ----------

func resolveCancelActivityType(orderType int64) int {
	switch orderType {
	case multi.ListingOrder:
		return multi.CancelListing
	case multi.CollectionBidOrder:
		return multi.CancelCollectionBid
	default:
		return multi.CancelItemBid
	}
}

// ---------- 业务构建与持久化层 ----------
func (s *Service) saveCancelActivity(ctx context.Context, cancelOrder multi.Order, log *ethereumTypes.Log) error {
	blockTime, err := s.chainClient.BlockTimeByNumber(ctx, big.NewInt(int64(log.BlockNumber)))
	if err != nil {
		return fmt.Errorf("get block time for block %d: %w", log.BlockNumber, err)
	}

	activity := multi.Activity{
		ActivityType:      resolveCancelActivityType(cancelOrder.OrderType),
		Maker:             cancelOrder.Maker,
		Taker:             ZeroAddress,
		MarketplaceID:     multi.MarketOrderBook,
		CollectionAddress: cancelOrder.CollectionAddress,
		TokenId:           cancelOrder.TokenId,
		CurrencyAddress:   s.ethAddress,
		Price:             cancelOrder.Price,
		BlockNumber:       int64(log.BlockNumber),
		TxHash:            log.TxHash.String(),
		EventTime:         int64(blockTime),
	}

	return s.db.WithContext(ctx).
		Table(multi.ActivityTableName(s.chain)).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&activity).Error
}

func (s *Service) pushCancelPriceEvent(cancelOrder multi.Order) {
	tradeEvent := &ordermanager.TradeEvent{
		OrderId:        cancelOrder.OrderID,
		CollectionAddr: cancelOrder.CollectionAddress,
		TokenID:        cancelOrder.TokenId,
		EventType:      ordermanager.Cancel,
	}

	if err := ordermanager.AddUpdatePriceEvent(s.kv, tradeEvent, s.chain); err != nil {
		slog.Error("failed to add update price event", "error", err, "type", "cancel", "order_id", cancelOrder.OrderID)
	}
}

func (s *Service) UpKeepingCollectionFloorChangeLoop(ctx context.Context) {
	// 1. 启动前校验索引状态
	var indexedStatus base.IndexedStatus
	if err := s.db.WithContext(ctx).Table(base.IndexedStatusTableName()).
		Select("last_indexed_time").
		Where("chain_id = ? and index_type = ?", s.chainId, CollectionFloorChangeIndexType).
		First(&indexedStatus).Error; err != nil {
		slog.ErrorContext(ctx, "failed to get collection floor change index status")
		return
	}
	// 若 last_indexed_time 为零值，说明尚未完成首次索引，跳过本次循环的初始化执行
	if indexedStatus.LastIndexedTime == 0 {
		slog.WarnContext(ctx, "collection floor change index not initialized yet, skip initial execution")
	} else {
		// 3. 服务启动后立即执行一次清理和更新，无需等待第一个 tick
		s.executeFloorChangeUpkeep(ctx)
	}

	timer := time.NewTicker(DaySeconds * time.Second)
	defer timer.Stop()
	updateFloorPriceTimer := time.NewTicker(MaxCollectionFloorTimeDifference * time.Second)
	defer updateFloorPriceTimer.Stop()

	for {
		select {
		case <-ctx.Done():
			slog.InfoContext(ctx, "UpKeepingCollectionFloorChangeLoop stopped due to context cancellation")
			return
		case <-timer.C:
			if err := s.deleteExpireCollectionFloorChangeFromDatabase(); err != nil {
				slog.ErrorContext(ctx, "failed to delete expire collection floor change")
			}
		case <-updateFloorPriceTimer.C:
			if s.projectName == gdb.OrderBookDexProject {
				floorPrices, err := s.QueryCollectionsFloorPrice(ctx)
				if err != nil {
					slog.ErrorContext(ctx, "failed to query collections floor change")
					continue
				}
				if err := s.persistCollectionsFloorChange(floorPrices); err != nil {
					slog.ErrorContext(ctx, "failed to persist collections floor price")
					continue
				}
			}
		}
	}
}

// executeFloorChangeUpkeep 封装启动时的立即执行逻辑，与定时器触发的逻辑保持一致
func (s *Service) executeFloorChangeUpkeep(ctx context.Context) {
	if err := s.deleteExpireCollectionFloorChangeFromDatabase(); err != nil {
		slog.ErrorContext(ctx, "failed to initial delete expire collection floor change")
	}

	if s.projectName == gdb.OrderBookDexProject {
		floorPrices, err := s.QueryCollectionsFloorPrice(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to initial query collections floor change")
			return
		}
		if err := s.persistCollectionsFloorChange(floorPrices); err != nil {
			slog.ErrorContext(ctx, "failed to initial persist collections floor price")
		}
	}
}

func (s *Service) deleteExpireCollectionFloorChangeFromDatabase() error {
	stmt := fmt.Sprintf(`DELETE FROM %s where event_time < UNIX_TIMESTAMP() - %d`, gdb.GetMultiProjectCollectionFloorPriceTableName(s.projectName, s.chain), CollectionFloorTimeRange)

	if err := s.db.Exec(stmt).Error; err != nil {
		return errors.Wrap(err, "failed to delete expire collection floor price")
	}

	return nil
}

func (s *Service) QueryCollectionsFloorPrice(ctx context.Context) ([]multi.CollectionFloorPrice, error) {
	timestamp := time.Now().Unix()
	timestampMilli := time.Now().UnixMilli()
	var collectionFloorPrice []multi.CollectionFloorPrice
	sql := fmt.Sprintf(`SELECT co.collection_address as collection_address,min(co.price) as price
FROM %s as ci
         left join %s co on co.collection_address = ci.collection_address and co.token_id = ci.token_id
WHERE (co.order_type = ? and
       co.order_status = ? and expire_time > ? and co.maker = ci.owner) group by co.collection_address`, gdb.GetMultiProjectItemTableName(s.projectName, s.chain), gdb.GetMultiProjectOrderTableName(s.projectName, s.chain))
	if err := s.db.WithContext(ctx).Raw(
		sql,
		multi.ListingType,
		multi.OrderStatusActive,
		time.Now().Unix(),
	).Scan(&collectionFloorPrice).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get collection floor price")
	}

	for i := 0; i < len(collectionFloorPrice); i++ {
		collectionFloorPrice[i].EventTime = timestamp
		collectionFloorPrice[i].CreateTime = timestampMilli
		collectionFloorPrice[i].UpdateTime = timestampMilli
	}

	return collectionFloorPrice, nil
}

func (s *Service) persistCollectionsFloorChange(FloorPrices []multi.CollectionFloorPrice) error {
	for i := 0; i < len(FloorPrices); i += DBBatchSizeLimit {
		end := i + DBBatchSizeLimit
		if i+DBBatchSizeLimit >= len(FloorPrices) {
			end = len(FloorPrices)
		}

		valueStrings := make([]string, 0)
		valueArgs := make([]interface{}, 0)

		for _, t := range FloorPrices[i:end] {
			valueStrings = append(valueStrings, "(?,?,?,?,?)")
			valueArgs = append(valueArgs, t.CollectionAddress, t.Price, t.EventTime, t.CreateTime, t.UpdateTime)
		}

		stmt := fmt.Sprintf(`INSERT INTO %s (collection_address,price,event_time,create_time,update_time)  VALUES %s
		ON DUPLICATE KEY UPDATE update_time=VALUES(update_time)`, gdb.GetMultiProjectCollectionFloorPriceTableName(s.projectName, s.chain), strings.Join(valueStrings, ","))

		if err := s.db.Exec(stmt, valueArgs...).Error; err != nil {
			return errors.Wrap(err, "failed to persist collection floor price info")
		}
	}
	return nil
}
