package storage

import (
	"url-shortener/internal/app/models"
)

type URLStorage interface {
	Get(shortURL string) (models.URLData, error)
	Add(url, shortURL, userID string) (string, error)
	Delete(url string) error
	AddBatch(listBatches []models.Batch, userID string) error
	GetUsersURLsByID(userID string) ([]models.URL, error)
	BatchUpdate(shortURLs []string, userID string) error
}
