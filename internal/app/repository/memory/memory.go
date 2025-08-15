// Пакет memory используется для записи и доступа данных в память.
package memory

import (
	"context"
	"strconv"
	"sync"

	config "url-shortener/internal/app/config/env"
	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/app/repository/model"
	"url-shortener/internal/pkg/constants"
	"url-shortener/internal/pkg/osfile"
)

// New создает новый Storage.
func New(cfg *config.Env) *Storage {
	return &Storage{
		mapURL: make(map[string]string),
		mu:     sync.RWMutex{},
		config: cfg,
	}
}

// GetOriginalByID возвращает оригинальный URL по сокращенному.
func (us *Storage) GetOriginalByID(context context.Context, url string) (model.URLData, error) {
	var URLData model.URLData
	value, found := us.mapURL[url]
	if !found {
		return URLData, constants.ErrURLNotFound
	}

	URLData.URL = value
	return URLData, nil
}

// GetByOriginalURL возвращает сокращенный URL по оригинальному.
func (us *Storage) GetByOriginalURL(context context.Context, originalURL string) (string, error) {
	return "", nil
}

// AddShorten добавляет пару оригинального и короткого URL в память.
func (us *Storage) AddShorten(context context.Context, url, shortURL, userID string) (string, error) {
	value, found := us.mapURL[url]
	if found {
		return value, constants.ErrURLAlreadyExists
	}

	us.mapURL[url] = shortURL
	us.mapURL[shortURL] = url

	return "", nil
}

// Write сохраняет событие в файл с использованием пакета osfile.
func (us *Storage) Write(url, shortURL string) error {
	data := osfile.Event{UUID: strconv.Itoa(len(us.mapURL)), ShortURL: shortURL, OriginalURL: url}
	consumer, err := osfile.NewConsumer(us.config.FileStoragePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = consumer.Close()
	}()

	if err = consumer.WriteEvent(&data); err != nil {
		return err
	}

	return nil
}

// AddBatch добавляет сразу несколько URL из батча в память и записывает их в файл.
func (us *Storage) AddBatch(context context.Context, batches []entity.Batch, userID string) error {
	for _, batch := range batches {
		_, err := us.AddShorten(context, batch.OriginalURL, batch.ShortURL, userID)
		if err != nil {
			return err
		}
		err = us.Write(batch.OriginalURL, batch.ShortURL)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetAllByUserID не реализован и возвращает ошибку ErrUnsupportedMethod.
func (us *Storage) GetAllByUserID(context context.Context, userID string) ([]model.URL, error) {
	return nil, constants.ErrUnsupportedMethod
}

// DeleteBatch не реализован и возвращает ошибку ErrUnsupportedMethod.
func (us *Storage) DeleteBatch(event entity.DeleteBatch) error {
	return constants.ErrUnsupportedMethod
}

func (us *Storage) GetStats(context context.Context) (int, int, error) {
	return 0, 0, constants.ErrUnsupportedMethod
}
