package postgres

import (
	url_repository "url-shortener/internal/app/domain/repository/url"
)

type Repositories struct {
	URL url_repository.IURLRepository
}
