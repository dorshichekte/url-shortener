package storage

import (
	"url-shortener/internal/app/models"
)

type URLStorage interface {
	Get(url string) (string, error)
	Add(url, shortURL string)
	Delete(url string) error
	AddBatch(listBatches []models.Batch) error
}
