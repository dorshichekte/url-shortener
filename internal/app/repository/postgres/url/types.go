package urlrepositorypostgres

import (
	"database/sql"

	config "url-shortener/internal/app/config/env"
)

type urlRepositoryPostgres struct {
	db     *sql.DB
	config *config.Env
}
