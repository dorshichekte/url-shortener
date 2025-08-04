package urlusecase

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
)

// GetOriginalByID возвращает оригинальный URL и статус удаления по короткому URL.
func (u *URLUseCase) GetOriginalByID(ctx context.Context, shortURL string) (entity.URLData, error) {
	URLData, err := u.URLRepository.GetOriginalByID(ctx, shortURL)
	if err != nil {
		return entity.URLData{}, err
	}

	urlData := entity.URLData{IsDeleted: URLData.IsDeleted, URL: URLData.URL}
	return urlData, nil
}
