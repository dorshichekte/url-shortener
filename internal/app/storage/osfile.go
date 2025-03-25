package storage

import (
	"encoding/json"
	"os"
)

func NewConsumer(fileName string) (*Consumer, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
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

func NewProducer(fileName string) (*Producer, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	return &Producer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (p *Producer) ReadEvent(s *URLStorage, fileStoragePath string) (*Event, error) {
	var event Event
	for p.decoder.Decode(&event) == nil {
		s.Add(event.OriginalURL, event.ShortURL, fileStoragePath)
	}
	return &event, nil
}

func (p *Producer) Close() error {
	return p.file.Close()
}
