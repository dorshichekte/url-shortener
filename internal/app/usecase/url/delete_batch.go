package urlusecase

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
)

func (u *URLUseCase) DeleteBatch(ctx context.Context, event entity.DeleteBatch) error {

	return u.URLRepository.DeleteBatch(ctx, event)
}
