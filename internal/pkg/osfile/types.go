package osfile

import (
	"encoding/json"
	"os"
)

// Consumer отвечает за запись событий в файл.
type Consumer struct {
	file    *os.File
	encoder *json.Encoder
}

// Event представляет структуру события, которое сохраняется в файл.
type Event struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// Producer отвечает за чтение событий из файла.
type Producer struct {
	file    *os.File
	decoder *json.Decoder
}
