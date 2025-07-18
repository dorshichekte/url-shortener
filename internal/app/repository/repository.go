package repository

import (
	"database/sql"

	c "url-shortener/internal/app/config/env"
	url_repository "url-shortener/internal/app/domain/repository/url"
	"url-shortener/internal/app/repository/memory"
	"url-shortener/internal/app/repository/postgres"
)

func New(db *sql.DB, config *c.Env) url_repository.IURLRepository {
	if config.DatabaseDSN == "" {
		return memory.New(config)
	}

	return postgres.New(db, config).URL
}
