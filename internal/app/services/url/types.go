package url

import (
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/models"
	"url-shortener/internal/app/storage"
)

type Methods interface {
	Shorten(originalURL, userID string) (string, error)
	GetOriginal(shortURL string) (models.URLData, error)
	BatchShorten(batch []models.BatchRequest) ([]models.BatchResponse, error)
	GetByUserID(userID string) ([]models.URL, error)
	BatchDelete(event models.DeleteEvent) error
}

type Service struct {
	Store storage.URLStorage
	common.BaseDependency
}
