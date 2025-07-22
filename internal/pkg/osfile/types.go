package osfile

import (
	"encoding/json"
	"os"
)

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
