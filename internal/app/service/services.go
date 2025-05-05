package service

import (
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/service/url"
	"url-shortener/internal/app/service/worker"
	"url-shortener/internal/app/storage"
)

func NewServices(storage storage.URLStorage, dependency common.BaseDependency) Services {
	return Services{
		URL:    url.NewService(storage, dependency),
		Worker: worker.NewService(20, 200),
	}
}
