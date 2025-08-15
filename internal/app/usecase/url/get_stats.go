package urlusecase

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
)

// GetStats возвращает количество сокращенных урл, и количество уникальных юзеров.
func (u *URLUseCase) GetStats(ctx context.Context) (entity.ServiceStats, error) {
	urlCount, userCount, err := u.URLRepository.GetStats(ctx)
	if err != nil {
		return entity.ServiceStats{}, err
	}

	data := entity.ServiceStats{URLCount: urlCount, UserCount: userCount}
	return data, nil
}
