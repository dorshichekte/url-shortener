// Пакет grpcadapter инициализирует сервер и роутер приложения, подключает мидлварины.
package grpcadapter

import (
	"context"

	"go.uber.org/zap"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/usecase"
	s "url-shortener/internal/pkg/server/grpc"
	"url-shortener/internal/pkg/validator"
)

// New создание экземпляра GRPCAdapter
func New(logger *zap.Logger, config *config.Config, cases *usecase.UseCases, validator *validator.Validator) *GRPCAdapter {
	server := s.New(logger, config, cases, validator)

	return &GRPCAdapter{
		server: server,
	}
}

// Start запуск grpc сервера
func (g *GRPCAdapter) Start(ctx context.Context) error {
	return g.server.Start(ctx)
}
