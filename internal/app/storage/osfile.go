package storage

import (
	"encoding/json"
	"io"
	"os"
)

func NewConsumer(filePath string) (*Consumer, error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
	file, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &Producer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}

func (p *Producer) ReadEvent(s *URLStorage) (*Event, error) {
	var event Event
	for {
		if err := p.decoder.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		s.Add(event.OriginalURL, event.ShortURL)
	}
	return &event, nil
}

func (p *Producer) Close() error {
	return p.file.Close()
}
