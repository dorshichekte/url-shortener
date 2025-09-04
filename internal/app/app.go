// Пакет app инициализирует зависимости приложения.
package app

import (
	"go.uber.org/zap"

	grpcadapter "url-shortener/internal/app/adapter/primary/grpc"
	hp "url-shortener/internal/app/adapter/primary/http"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/repository"
	pg "url-shortener/internal/app/repository/postgres"
	"url-shortener/internal/app/usecase"
	a "url-shortener/internal/pkg/auth"
	v "url-shortener/internal/pkg/validator"
)

// New создает и инициализирует все зависимости приложения: логгер, конфиг, базу данных,
func New(logger *zap.Logger, config *config.Config) *App {
	validator := v.New()
	auth := a.New(config.Env.AccessSecretKey)

	postgresConnection := pg.NewConnection(logger, config.Env)

	repositories := repository.New(postgresConnection, config.Env)

	useCases := usecase.New(config.Env, repositories, postgresConnection)

	grpcAdapter := grpcadapter.New(logger, config, useCases, validator)
	httpAdapter := hp.New(logger, auth, config, useCases, validator)

	return &App{
		HTTPAdapter: httpAdapter,
		GRPCAdapter: grpcAdapter,
	}
}
