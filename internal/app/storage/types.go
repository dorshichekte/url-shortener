package storage

import (
	"url-shortener/internal/app/models"
)

type URLStorage interface {
	Get(url string) (string, error)
	Add(url, shortURL, userID string)
	Delete(url string) error
	AddBatch(listBatches []models.Batch, userID string) error
	GetUsersURLsByID(userID string) ([]models.URL, error)
}
