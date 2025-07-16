package config

import (
	adapter "url-shortener/internal/app/config/adapter"
	"url-shortener/internal/app/config/env"
)

type Config struct {
	Env         *config.Env
	HTTPAdapter *adapter.HTTPAdapter
}
