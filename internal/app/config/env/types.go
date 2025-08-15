package config

// Env конфигурация env конфига приложения.
type Env struct {
	ServerAddress   string `env:"SERVER_ADDRESS" json:"server_address"`
	DatabaseDSN     string `env:"DATABASE_DSN" json:"database_dsn"`
	FileStoragePath string `env:"FILE_STORAGE_PATH" json:"file_storage_path"`
	BaseURL         string `env:"BASE_URL" json:"base_url"`
	AccessSecretKey string `env:"ACCESS_SECRET_KEY" json:"access_secret_key"`
	EnableHTTPS     bool   `env:"ENABLE_HTTPS" json:"enable_https"`
	Config          string `env:"CONFIG"`
	TrustedSubnet   string `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
}
