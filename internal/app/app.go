package app

import (
	"context"

	"go.uber.org/zap"

	hp "url-shortener/internal/app/adapter/primary/http"
	"url-shortener/internal/app/config"
	pg "url-shortener/internal/app/repository/postgres"
	"url-shortener/internal/app/usecase"
	a "url-shortener/internal/pkg/auth"
	v "url-shortener/internal/pkg/validator"
)

func New(ctx context.Context, logger *zap.Logger, config *config.Config) *App {
	validator := v.New()
	auth := a.New(config.Env.AccessSecretKey)

	postgresConnection := pg.NewConnection(logger, config.Env).Db

	repositories := pg.New(postgresConnection, config.Env)

	useCases := usecase.New(config.Env, repositories)

	httpAdapter := hp.New(logger, auth, config, useCases, validator, postgresConnection)

	return &App{
		HTTPAdapter: httpAdapter,
	}
}
