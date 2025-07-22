package urlusecase

import (
	entity "url-shortener/internal/app/domain/entity/url"
)

func (u *URLUseCase) DeleteBatch(event entity.DeleteBatch) error {
	return u.URLRepository.DeleteBatch(event)
}
