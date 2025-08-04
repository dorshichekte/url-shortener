// Пакет config инициализирует конфиг приложения.
package config

import (
	"github.com/joho/godotenv"

	adapter "url-shortener/internal/app/config/adapter"
	env "url-shortener/internal/app/config/env"
	worker "url-shortener/internal/app/config/worker"
)

// New создает экземпляр Config.
func New() (*Config, error) {
	_ = godotenv.Load()
	envConfig, err := env.New()
	if err != nil {
		return &Config{}, err
	}

	httpAdapterConfig := adapter.New(envConfig.ServerAddress)

	workerConfig := worker.New()

	return &Config{
		Env:         envConfig,
		HTTPAdapter: httpAdapterConfig,
		Worker:      workerConfig,
	}, nil
}
