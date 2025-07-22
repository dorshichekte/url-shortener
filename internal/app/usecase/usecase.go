package usecase

import (
	config "url-shortener/internal/app/config/env"
	url_repository "url-shortener/internal/app/domain/repository/url"
	url "url-shortener/internal/app/usecase/url"
)

func New(config *config.Env, repositories url_repository.IURLRepository) *UseCases {
	return &UseCases{
		URL: url.New(config, repositories),
	}
}
