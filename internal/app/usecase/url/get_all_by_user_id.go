package urlusecase

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
)

func (u *URLUseCase) GetAllByUserID(ctx context.Context, userID string) ([]entity.URL, error) {
	_, err := u.URLRepository.GetAllByUserID(ctx, userID)
	return nil, err
}
