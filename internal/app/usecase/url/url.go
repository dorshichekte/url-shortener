// Пакет urlusecase включает бизнес логику для сущности урл.
package urlusecase

import (
	config "url-shortener/internal/app/config/env"
	url_repository "url-shortener/internal/app/domain/repository/url"
)

func New(config *config.Env, urlRepository url_repository.IURLRepository) *URLUseCase {
	return &URLUseCase{
		URLRepository: urlRepository,
		Config:        config,
	}
}
