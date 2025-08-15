// Пакет urlrepository включает интерфейс с описанием сигнатуры методов для URL-репозитория.
package urlrepository

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/app/repository/model"
)

// IURLRepository описывает поведение репозитория для работы с URL.
//
//go:generate mockgen -package mock -source url.go -destination ../../mock/url_repositoy.go UrlRepostiry
type IURLRepository interface {
	AddShorten(context context.Context, originalURL, shortURL, userID string) (string, error)
	GetOriginalByID(context context.Context, shortURL string) (model.URLData, error)
	GetByOriginalURL(context context.Context, originalURL string) (string, error)
	GetAllByUserID(context context.Context, userID string) ([]model.URL, error)
	AddBatch(context context.Context, batches []entity.Batch, userID string) error
	DeleteBatch(event entity.DeleteBatch) error
	GetStats(context context.Context) (int, int, error)
}
