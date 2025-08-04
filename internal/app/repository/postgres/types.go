package postgres

import (
	url_repository "url-shortener/internal/app/domain/repository/url"
)

// Repositories агрегирует все репозитории, которые используют PostgreSQL в качестве хранилища.
type Repositories struct {
	URL url_repository.IURLRepository
}
