package config

import (
	adapter "url-shortener/internal/app/config/adapter"
	config "url-shortener/internal/app/config/env"
	worker "url-shortener/internal/app/config/worker"
)

type Config struct {
	Env         *config.Env
	HTTPAdapter *adapter.HTTPAdapter
	Worker      *worker.Worker
}
