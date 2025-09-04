// Пакет httpadapter инициализирует сервер и роутер приложения, подключает мидлварины.
package httpadapter

import (
	"context"
	"net/http"
	httpserver "url-shortener/internal/pkg/server/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/handler"
	"url-shortener/internal/app/adapter/primary/http/router"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/usecase"
	"url-shortener/internal/pkg/auth"
	"url-shortener/internal/pkg/validator"
)

// New создает новый экземпляр HTTPAdapter.
func New(logger *zap.Logger, auth auth.Auth, config *config.Config, useCases *usecase.UseCases, validator *validator.Validator) *HTTPAdapter {
	rtr := newRouter(logger, auth, config, useCases, validator)

	s := httpserver.New(logger, config, rtr)

	return &HTTPAdapter{
		server: s,
	}
}

func newRouter(logger *zap.Logger, auth auth.Auth, config *config.Config, useCases *usecase.UseCases, validator *validator.Validator) http.Handler {
	r := router.New(logger)

	h := handler.New(logger, config.Env, useCases, validator)

	r.AppendRoutes(config, h, auth)

	return r.Router()
}

// Start запускает HTTP-сервер.
func (a *HTTPAdapter) Start(ctx context.Context) error {
	return a.server.Start(ctx)
}
