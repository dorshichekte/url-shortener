package config

// Env конфигурация env конфига приложения.
type Env struct {
	ServerAddress   string `env:"SERVER_ADDRESS"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	BaseURL         string `env:"BASE_URL"`
	AccessSecretKey string `env:"ACCESS_SECRET_KEY"`
}
