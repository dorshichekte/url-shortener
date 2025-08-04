package urlusecase

import (
	entity "url-shortener/internal/app/domain/entity/url"
)

// DeleteBatch отмечает пакет URL на удаление для указанного пользователя.
func (u *URLUseCase) DeleteBatch(event entity.DeleteBatch) error {
	return u.URLRepository.DeleteBatch(event)
}
