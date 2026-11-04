package service

import (
	"MiSwapSync/chain/chainclient"
	"MiSwapSync/config"
	"MiSwapSync/ordermanager"
	"MiSwapSync/service/collectionfilter"
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
	//orderbookSyncer  *orderbookindexer.Service	//数据解析
	wg sync.WaitGroup
}

// New TODO 新建service
//
//	func New(ctx context.Context, cfg *config.Config) (*Service, error) {
//		var kvConf kv.KvConf
//		//遍历redis的config，拼接切片信息，拿到redis的配置信息
//		for _, conn := range cfg.Kv.Redis {
//			kvConf = append(kvConf, cache.NodeConf{RedisConf: redis.RedisConf{
//				Host: conn.Host,
//				Type: conn.Type,
//				Pass: conn.Pass,
//			},
//				Weight: 2,
//			})
//		}
//		//获取redis连接信息
//		kvStore := xkv.NewStore(kvConf)
//
//		//获取数据库
//		db := gdb.MustNew(cfg.DB)
//
//		manager := Service{
//			kvStore: kvStore,
//			db:      db,
//		}
//
//		return &manager, nil
//	}
/*
基础设施（DB、Redis、ChainClient）→ 外部传入（因为它们是跨层共享资源，且生命周期由 run.go 管理）
业务组件（Filter、Manager、Syncer）→ New() 内部创建（因为它们是 Service 的私有实现细节）
*/
func New(ctx context.Context, cfg *config.Config, db *gorm.DB, kvStore *xkv.Store, chainClient chainclient.ChainClient) (*Service, error) {
	//筛选器
	filter := collectionfilter.New(db, cfg.ChainCfg.Name, cfg.ProjectCfg.Name)
	//订单管理器
	orderManager := ordermanager.NewOrderManager(db, cfg.ChainCfg.Name, cfg.ProjectCfg.Name)
	//TODO 服务解析器
	//TODO 订单管理服务

	return &Service{
		db:               db,
		kvStore:          kvStore,
		collectionFilter: filter,
		orderManager:     orderManager,
	}, nil
}

func (s *Service) Start(ctx context.Context) error {
	if err := s.collectionFilter.PreloadCollections(ctx); err != nil {
		return errors.Wrap(err, "failed to preload collections")
	}
	s.wg.Add(2)
	go func() {
		defer s.wg.Done()
		//TODO orderbookIndex服务启动,内部监听 ctx.Done()
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
