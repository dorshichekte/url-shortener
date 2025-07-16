package repository

import (
	"sync"

	"url-shortener/internal/app/repository/model"
	"url-shortener/internal/config"
)

type URLStorage interface {
	Get(shortURL string) (model.URLData, error)
	Add(url, shortURL, userID string) (string, error)
	Delete(url string) error
	AddBatch(listBatches []model.Batch, userID string) error
	GetURLsByID(userID string) ([]model.URL, error)
	BatchUpdate(event model.DeleteEvent) error
}

type BaseStorageDependency struct {
	cfg config
	mu  sync.RWMutex
}
