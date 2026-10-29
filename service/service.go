package service

import (
	"MiSwapSync/config"
	"context"
)

type Service struct {
	config *config.Config
}

// New TODO 新建service
func New(ctx context.Context, cfg *config.Config) (*Service, error) {

	return nil, nil
}

// Start TODO 启动服务
func (s *Service) Start() error {

	return nil
}
