package osfile

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func NewConsumer(filePath string) (*Consumer, error) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	fmt.Println(filePath, "filepath")
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

func (c *Consumer) Load(fileStoragePath string) (*[]Event, error) {
	pr, err := NewProducer(fileStoragePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = pr.Close()
	}()

	listEvents, err := pr.ReadEvent()
	if err != nil {
		return nil, err
	}

	return listEvents, nil
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

func (p *Producer) ReadEvent() (*[]Event, error) {
	var events []Event
	for {
		var event Event
		if err := p.decoder.Decode(&event); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		events = append(events, event)
	}

	return &events, nil
}

func (p *Producer) Close() error {
	return p.file.Close()
}
