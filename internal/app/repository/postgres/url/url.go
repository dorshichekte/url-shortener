// Пакет urlrepositorypostgres вклчюает методы для работы с урл репозиторием.
package urlrepositorypostgres

import (
	"database/sql"

	config "url-shortener/internal/app/config/env"
	urlrepository "url-shortener/internal/app/domain/repository/url"
)

// New создает и возвращает новый экземпляр urlRepositoryPostgres,
func New(db *sql.DB, config *config.Env) urlrepository.IURLRepository {
	return &urlRepositoryPostgres{db: db, config: config}
}
