package storage

import (
	"encoding/json"
	"os"
)

type URLType string

type MapURL map[string]string

type URLStorage struct {
	mapURL      MapURL
	mapShortURL MapURL
}

type Consumer struct {
	file    *os.File
	encoder *json.Encoder
}

type Event struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type Producer struct {
	file    *os.File
	decoder *json.Decoder
}
