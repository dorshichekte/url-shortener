package urlusecase

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
)

func (u *URLUseCase) GetAllByUserID(ctx context.Context, userID string) ([]entity.URL, error) {
	urls, err := u.URLRepository.GetAllByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var urlData []entity.URL
	for _, url := range urls {
		urlData = append(urlData, entity.URL{ShortURL: url.ShortURL, OriginalURL: url.OriginalURL})
	}

	return urlData, err
}
