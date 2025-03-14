package storage

type URLType string

type MapURL map[string]string

type URLStorage struct {
	mapURL      MapURL
	mapShortURL MapURL
}
