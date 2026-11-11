package service

import (
	"MiSwapSync/chain"
	"MiSwapSync/chain/chainclient"
	"MiSwapSync/config"
	"MiSwapSync/ordermanager"
	"MiSwapSync/service/collectionfilter"
	"MiSwapSync/service/orderbookindex"
	"MiSwapSync/stores/xkv"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	kvStore          *xkv.Store
	db               *gorm.DB
	collectionFilter *collectionfilter.Filter   //筛选器
	orderManager     *ordermanager.OrderManager //订单管理
	orderbookSyncer  *orderbookindex.Service    //数据解析
	wg               sync.WaitGroup
}

/*
基础设施（DB、Redis、ChainClient）→ 外部传入（因为它们是跨层共享资源，且生命周期由 run.go 管理）
业务组件（Filter、Manager、Syncer）→ New() 内部创建（因为它们是 Service 的私有实现细节）
*/

// New 新建
func New(ctx context.Context, cfg *config.Config, db *gorm.DB, kvStore *xkv.Store, chainClient chainclient.ChainClient) (*Service, error) {
	//筛选器
	filter := collectionfilter.New(db, cfg.ChainCfg.Name, cfg.ProjectCfg.Name)
	//订单管理器
	orderManager := ordermanager.NewOrderManager(db, cfg.ChainCfg.Name, cfg.ProjectCfg.Name)
	var orderBookSyncer *orderbookindex.Service
	var err error
	chainClient, err = chainclient.New(int(cfg.ChainCfg.ID), cfg.AnkrCfg.HttpsUrl+cfg.AnkrCfg.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create evm client")
	}

	switch cfg.ChainCfg.ID {
	case chain.EthChainID, chain.OptimismChainID, chain.SepoliaChainID:
		//参数对应的值 11155111 sepolia
		orderBookSyncer = orderbookindex.New(ctx, cfg, db, kvStore, chainClient, cfg.ChainCfg.ID, cfg.ChainCfg.Name, orderManager)
	}
	return &Service{
		db:               db,
		kvStore:          kvStore,
		collectionFilter: filter,
		orderManager:     orderManager,
		orderbookSyncer:  orderBookSyncer,
	}, nil
}

func (s *Service) Start(ctx context.Context) error {
	if err := s.collectionFilter.PreloadCollections(ctx); err != nil {
		return errors.Wrap(err, "failed to preload collections")
	}
	s.wg.Add(2)
	go func() {
		defer s.wg.Done()
		s.orderbookSyncer.Start(ctx)
	}()
	go func() {
		defer s.wg.Done()
		s.orderManager.Start(ctx)
	}()
	//s.wg.Wait()
	return nil //← 非阻塞返回，把控制权交还给调用方
}

// WaitShutdown 供外部包等待所有子服务安全退出
func (s *Service) WaitShutdown() {
	s.wg.Wait()
}
