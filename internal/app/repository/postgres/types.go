package postgres

import (
	"database/sql"

	url_repository "url-shortener/internal/app/domain/repository/url"
)

type Postgres struct {
	DB *sql.DB
}

type Repositories struct {
	URL url_repository.IURLRepository
}
