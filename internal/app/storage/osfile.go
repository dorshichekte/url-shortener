package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const fileName = "url-db.json"

func NewConsumer(filePath string) (*Consumer, error) {
	if filePath[len(filePath)-1] != '/' {
		filePath += "/"
	}

	file, err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (c *Consumer) WriteEvent(data *Event) error {
	return c.encoder.Encode(&data)
}

func (c *Consumer) Close() error {
	return c.file.Close()
}

func NewProducer(filePath string) (*Producer, error) {
	createFile(filePath)

	if filePath[len(filePath)-1] != '/' {
		filePath += "/"
	}

	file, err := os.OpenFile(filePath+fileName, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	log.Printf("Файл для хранения данных URL создан по пути: %s", filePath+fileName)

	return &Producer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (p *Producer) ReadEvent(s *URLStorage) (*Event, error) {
	var event Event
	for p.decoder.Decode(&event) == nil {
		fmt.Println(event)
		s.Add(event.OriginalURL, event.ShortURL)
	}
	return &event, nil
}

func (p *Producer) Close() error {
	return p.file.Close()
}

func createFile(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filePath, 0755)
	}
}
