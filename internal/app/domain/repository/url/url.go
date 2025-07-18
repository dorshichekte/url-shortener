package urlrepository

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/app/repository/model"
)

type IURLRepository interface {
	AddShorten(context context.Context, originalURL, shortUrl, userID string) (string, error)
	GetOriginalByID(context context.Context, shortURL string) (model.URLData, error)
	GetAllByUserID(context context.Context, userID string) ([]model.URL, error)
	AddBatch(context context.Context, batches []entity.Batch, userID string) error
	DeleteBatch(event entity.DeleteBatch) error
}
