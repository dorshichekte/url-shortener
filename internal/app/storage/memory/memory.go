package memory

import (
	"strconv"
	"url-shortener/internal/app/config"

	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/models"
	"url-shortener/internal/app/osfile"
)

func NewURLStorage(cfg *config.Config) *Storage {
	return &Storage{
		mapURL: make(map[string]string),
		cfg:    cfg,
	}
}

func (us *Storage) Get(url string) (string, error) {
	value, found := us.mapURL[url]
	if !found {
		return "", constants.ErrURLNotFound
	}
	return value, nil
}

func (us *Storage) Add(url, shortURL string) {
	us.mapURL[url] = shortURL
	us.mapURL[shortURL] = url
}

func (us *Storage) Delete(url string) error {
	_, found := us.mapURL[url]
	if !found {
		return constants.ErrURLNotFound
	}

	delete(us.mapURL, url)
	return nil
}

func (us *Storage) Write(url, shortURL string) error {
	data := osfile.Event{UUID: strconv.Itoa(len(us.mapURL)), ShortURL: shortURL, OriginalURL: url}

	consumer, err := osfile.NewConsumer(us.cfg.FileStoragePath)
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

func (us *Storage) AddBatch(listBatches []models.Batch) error {
	for _, batch := range listBatches {
		us.Add(batch.OriginalURL, batch.ShortURL)
		err := us.Write(batch.OriginalURL, batch.ShortURL)
		if err != nil {
			return err
		}
	}

	return nil
}
