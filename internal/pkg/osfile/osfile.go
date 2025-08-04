// Пакет osfile предоставляет методы для работы с файлом.
package osfile

import (
	"encoding/json"
	"io"
	"os"
)

// NewConsumer открывает файл для записи (дописывания) и возвращает Consumer.
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

// WriteEvent записывает одно событие в файл в формате JSON.
func (c *Consumer) WriteEvent(data *Event) error {
	return c.encoder.Encode(&data)
}

// Close закрывает файл, открытый Consumer'ом.
func (c *Consumer) Close() error {
	return c.file.Close()
}

// Load считывает все события из файла, открытого через Producer.
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

// NewProducer открывает файл для чтения и возвращает Producer.
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

// ReadEvent читает все события из файла и возвращает их слайс.
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

// Close закрывает файл, открытый Producer'ом.
func (p *Producer) Close() error {
	return p.file.Close()
}
