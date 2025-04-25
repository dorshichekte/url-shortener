package memory

import (
	"sync"

	"url-shortener/internal/app/config"
)

type URLType string

type MapURL map[string]string

type Storage struct {
	mapURL MapURL
	mu     sync.Mutex
	cfg    config.Config
}
