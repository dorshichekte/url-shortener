package services

import (
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/services/url"
	"url-shortener/internal/app/storage"
)

func NewServices(storage storage.URLStorage, dependency common.BaseDependency) Services {
	return Services{
		URL: url.NewService(storage, dependency),
	}
}
