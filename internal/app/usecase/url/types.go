package urlusecase

import (
	"context"

	config "url-shortener/internal/app/config/env"
	entity "url-shortener/internal/app/domain/entity/url"
	url_repository "url-shortener/internal/app/domain/repository/url"
)

// IUrlUseCase описывает интерфейс бизнес-логики для работы с URL.
type IUrlUseCase interface {
	AddShorten(context context.Context, originalURL, userID string) (string, error)
	GetOriginalByID(context context.Context, shortURL string) (entity.URLData, error)
	GetAllByUserID(context context.Context, userID string) ([]entity.URL, error)
	AddBatch(context context.Context, batches []entity.Batch, userID string) ([]entity.Batch, error)
	DeleteBatch(event entity.DeleteBatch) error
}

// URLUseCase реализует бизнес-логику URL сервиса.
type URLUseCase struct {
	URLRepository url_repository.IURLRepository
	Config        *config.Env
}
