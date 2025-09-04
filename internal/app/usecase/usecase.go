// Пакет usecase инициализирует usercase приложения.
package usecase

import (
	"database/sql"

	config "url-shortener/internal/app/config/env"
	url_repository "url-shortener/internal/app/domain/repository/url"
	db "url-shortener/internal/app/usecase/db"
	url "url-shortener/internal/app/usecase/url"
)

// New создает и возвращает структуру UseCases с инициализированными бизнес-логиками.
func New(config *config.Env, repositories url_repository.IURLRepository, dbConnection *sql.DB) *UseCases {
	return &UseCases{
		DB:  db.New(dbConnection),
		URL: url.New(config, repositories),
	}
}
