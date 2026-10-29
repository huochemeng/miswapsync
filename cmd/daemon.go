package cmd

import (
	"MiSwapSync/config"
	"MiSwapSync/service"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
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
	// 启动核心业务服务
	s, err := service.New(ctx, cfg)
	if err != nil {
		slog.Error("Failed to create sync server", "error", err)
		onSyncExit <- err
		return
	}

	if err := s.Start(); err != nil {
		slog.Error("Failed to start sync server", "error", err)
		onSyncExit <- err
	}
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
