package url

import "url-shortener/internal/app/storage"

type Service struct {
	store storage.URLStorage
}
