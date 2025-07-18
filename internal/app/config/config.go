package config

import (
	"github.com/joho/godotenv"

	adapter "url-shortener/internal/app/config/adapter"
	env "url-shortener/internal/app/config/env"
)

func New() (*Config, error) {
	_ = godotenv.Load()
	envConfig, err := env.New()
	if err != nil {
		return &Config{}, err
	}

	httpAdapterConfig := adapter.New(envConfig.ServerAddress)

	return &Config{
		Env:         envConfig,
		HTTPAdapter: httpAdapterConfig,
	}, nil
}
