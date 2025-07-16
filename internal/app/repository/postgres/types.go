package postgres

import (
	"database/sql"

	url_repository "url-shortener/internal/app/domain/repository/url"
)

type Postgres struct {
	Db *sql.DB
}

type Repositories struct {
	Url url_repository.IURLRepository
}
