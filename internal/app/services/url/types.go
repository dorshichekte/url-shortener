package url

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/storage"
)

type Service struct {
	store storage.URLStorage
	cfg   config.Config
}
