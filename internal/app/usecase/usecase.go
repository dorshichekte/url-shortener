package usecase

import (
	config "url-shortener/internal/app/config/env"
	"url-shortener/internal/app/repository/postgres"
	url "url-shortener/internal/app/usecase/url"
)

func New(config *config.Env, repositories postgres.Repositories) *UseCases {
	return &UseCases{
		URL: url.New(config, repositories.URL),
	}
}
