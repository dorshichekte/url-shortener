package url

import (
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/model"
	"url-shortener/internal/app/storage"
)

type Methods interface {
	Shorten(originalURL, userID string) (string, error)
	GetOriginal(shortURL string) (model.URLData, error)
	BatchShorten(batch []model.BatchRequest) ([]model.BatchResponse, error)
	GetByUserID(userID string) ([]model.URL, error)
	BatchDelete(event model.DeleteEvent) error
}

type Service struct {
	Store storage.URLStorage
	common.BaseDependency
}
