package config

import (
	"github.com/joho/godotenv"

	adapter "url-shortener/internal/app/config/adapter"
	env "url-shortener/internal/app/config/env"
)

func New() (*Config, error) {
	_ = godotenv.Load()

	envCfg, err := env.New()
	if err != nil {
		return &Config{}, err
	}

	httpAdapterCfg := adapter.New(envCfg.ServerAddress)

	return &Config{
		Env:         envCfg,
		HTTPAdapter: httpAdapterCfg,
	}, nil
}
