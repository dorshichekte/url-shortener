package config

import "go.uber.org/zap"

type AppConfig struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	BaseURL         string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

type Config struct {
	App    AppConfig
	Logger *zap.Logger
}
