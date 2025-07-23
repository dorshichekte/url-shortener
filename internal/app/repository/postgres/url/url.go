// Package urlrepositorypostgres contains methods for working with url entity.
package urlrepositorypostgres

import (
	"database/sql"

	config "url-shortener/internal/app/config/env"
	urlrepository "url-shortener/internal/app/domain/repository/url"
)

func New(db *sql.DB, config *config.Env) urlrepository.IURLRepository {
	return &urlRepositoryPostgres{db: db, config: config}
}
