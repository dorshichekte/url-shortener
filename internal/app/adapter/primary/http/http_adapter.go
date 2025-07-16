package httpadapter

import (
	"context"
	"database/sql"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/handler"
	"url-shortener/internal/app/adapter/primary/http/router"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/usecase"
	"url-shortener/internal/pkg/auth"
	"url-shortener/internal/pkg/server"
	"url-shortener/internal/pkg/validator"
)

func New(logger *zap.Logger, auth auth.Auth, config *config.Config, useCases *usecase.UseCases, validator *validator.Validator, dbConnection *sql.DB) *HTTPAdapter {
	rtr := newRouter(logger, auth, config, useCases, validator, dbConnection)

	s := server.New(logger, &config.HTTPAdapter.Server, rtr)

	return &HTTPAdapter{
		server: s,
	}
}

func newRouter(logger *zap.Logger, auth auth.Auth, config *config.Config, useCases *usecase.UseCases, validator *validator.Validator, dbConnection *sql.DB) http.Handler {
	r := router.New(logger)

	h := handler.New(logger, config.Env, useCases, validator, dbConnection)

	r.AppendRoutes(&config.HTTPAdapter.Router, h, auth)

	return r.Router()
}

func (a HTTPAdapter) Start(ctx context.Context) error {
	return a.server.Start(ctx)
}
