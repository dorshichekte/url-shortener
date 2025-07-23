package memory

import (
	"sync"

	config "url-shortener/internal/app/config/env"
)

// URLType представляет тип для хранения URL в виде строки.
type URLType string

// MapURL определяет тип данных — отображение строк на строки.
type MapURL map[string]string

// Storage хранит URL-данные в памяти с обеспечением потокобезопасности.
type Storage struct {
	mapURL MapURL
	mu     sync.RWMutex
	config *config.Env
}
