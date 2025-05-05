package storage

import (
	"sync"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/models"
)

type URLStorage interface {
	Get(shortURL string) (models.URLData, error)
	Add(url, shortURL, userID string) (string, error)
	Delete(url string) error
	AddBatch(listBatches []models.Batch, userID string) error
	GetUsersURLsByID(userID string) ([]models.URL, error)
	BatchUpdate(event models.DeleteEvent) error
}

type BaseStorageDependency struct {
	cfg config.AppConfig
	mu  sync.RWMutex
}
