// Пакет handler инициализируют обработчики приложения.
package handler

import (
	"database/sql"

	"go.uber.org/zap"

	db_handler "url-shortener/internal/app/adapter/primary/http/handler/db"
	url_handler "url-shortener/internal/app/adapter/primary/http/handler/url"
	config "url-shortener/internal/app/config/env"
	"url-shortener/internal/app/usecase"
	"url-shortener/internal/pkg/validator"
)

func New(logger *zap.Logger, config *config.Env, useCases *usecase.UseCases, validator *validator.Validator, dbConnection *sql.DB) *Handlers {
	return &Handlers{
		Database: db_handler.New(logger, dbConnection),
		URL:      url_handler.New(logger, config, useCases.URL, validator),
	}
}
