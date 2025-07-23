// Пакет usecase инициализирует usercase приложения.
package usecase

import (
	config "url-shortener/internal/app/config/env"
	url_repository "url-shortener/internal/app/domain/repository/url"
	url "url-shortener/internal/app/usecase/url"
)

// New создает и возвращает структуру UseCases с инициализированными бизнес-логиками.
func New(config *config.Env, repositories url_repository.IURLRepository) *UseCases {
	return &UseCases{
		URL: url.New(config, repositories),
	}
}
