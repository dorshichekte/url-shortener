package osfile

import (
	"bufio"
	"os"
)

type Consumer struct {
	file    *os.File
	scanner *bufio.Scanner
}

type Event struct {
	UUID        string `json:"uuid"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type Producer struct {
	file   *os.File
	writer *bufio.Writer
}
