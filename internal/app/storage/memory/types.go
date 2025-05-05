package memory

import "url-shortener/internal/app/config"

type URLType string

type MapURL map[string]string

type Storage struct {
	mapURL MapURL
	cfg    config.Config
}
