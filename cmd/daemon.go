package cmd

import (
	"MiSwapSync/chain/chainclient"
	"MiSwapSync/config"
	"MiSwapSync/service"
	"MiSwapSync/stores/gdb"
	"MiSwapSync/stores/xkv"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// initSyncServer 封装所有初始化逻辑
func initSyncServer(ctx context.Context) (*config.Config, error) {
	//解析配置
	cfg, err := config.UnmarshalCmdConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall config: %w", err)
	}
	slog.Info("Sync server config loaded successfully", "config", cfg)
	return cfg, nil
}

// 封装服务的实际运行逻辑
func runSyncServer(ctx context.Context, cfg *config.Config, onSyncExit chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	slog.Info("Sync server starting ...")
	//启动pprof（独立 Goroutine，且响应 ctx 优雅停机）
	if cfg.Monitor.PprofEnable {
		srv := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%d", cfg.Monitor.PprofPort)}

		// 监听取消信号，优雅关闭 pprof
		go func() {
			<-ctx.Done()
			slog.Info("Shutting down pprof server...")
			if err := srv.Shutdown(context.Background()); err != nil {
				slog.Error("Pprof server shutdown failed", "error", err)
			}
		}()

		// 独立启动 pprof，避免阻塞当前协程
		go func() {
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				slog.Error("pprof server error", "error", err)
			}
		}()
	}

	// 2. 初始化基础设施（在 goroutine 内同步执行，失败直接上报）
	db, err := gdb.New(cfg.DB)
	if err != nil {
		slog.Error("Failed to init database", "error", err)
		onSyncExit <- err
		return
	}

	var kvConf kv.KvConf
	//遍历redis的config，拼接切片信息，拿到redis的配置信息
	for _, conn := range cfg.Kv.Redis {
		kvConf = append(kvConf, cache.NodeConf{RedisConf: redis.RedisConf{
			Host: conn.Host,
			Type: conn.Type,
			Pass: conn.Pass,
		},
			Weight: 2,
		})
	}
	//获取redis连接信息
	kvStore := xkv.NewStore(kvConf)

	chainClient, err := chainclient.New(
		int(cfg.ChainCfg.ID),
		cfg.AnkrCfg.HttpsUrl+cfg.AnkrCfg.ApiKey,
	)
	if err != nil {
		slog.Error("Failed to init chain client", "error", err)
		onSyncExit <- err
		return
	}
	// 启动核心业务服务
	s, err := service.New(ctx, cfg, db, kvStore, chainClient)
	if err != nil {
		slog.Error("Failed to create sync server", "error", err)
		onSyncExit <- err
		return
	}

	if err := s.Start(ctx); err != nil {
		slog.Error("Failed to start sync server", "error", err)
		onSyncExit <- err
		return
	}

	<-ctx.Done() // ← 保活：响应取消信号后才进入清理阶段
	slog.Info("Sync server shutting down, waiting for services...")
	s.WaitShutdown() // ← 在这里 Wait：确保所有子任务安全退出后再释放资源
	slog.Info("All services gracefully stopped")
}

var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "sync mi swap order info",
	Long:  "sync mi swap order info.",
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Daemon run")
		//wg定义在Run内容，更推荐的做法
		var wg sync.WaitGroup
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// 优雅退出通道
		onSyncExit := make(chan error, 1)
		onSignal := make(chan os.Signal, 1)
		signal.Notify(onSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		// 1. 同步执行初始化（不放进 Goroutine）
		cfg, err := initSyncServer(ctx)
		if err != nil {
			slog.Error("Initialization failed", "error", err)
			onSyncExit <- err
		} else {
			// 2. 初始化成功，启动服务
			wg.Add(1)
			go runSyncServer(ctx, cfg, onSyncExit, &wg)
		}

		// 3. 阻塞监听：等待系统信号或内部致命错误
		select {
		case sig := <-onSignal:
			cancel()
			slog.Info("Exit by signal", "signal", sig.String())
		case err := <-onSyncExit:
			cancel()
			slog.Error("Exit by error", "error", err)
		}

		// 4. 等待所有 Goroutine 安全退出
		wg.Wait()
		slog.Info("Server gracefully stopped.")
	},
}

func init() {
	slog.Info("daemon init")
	//子命令DaemonCmd添加到root命令
	rootCmd.AddCommand(DaemonCmd)
}
