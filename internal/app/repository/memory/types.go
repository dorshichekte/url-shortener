package memory

import (
	"sync"

	config "url-shortener/internal/app/config/env"
)

type URLType string

type MapURL map[string]string

type Storage struct {
	mapURL MapURL
	mu     sync.RWMutex
	config *config.Env
}
