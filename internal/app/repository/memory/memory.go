package memory

import (
	"strconv"
	"sync"

	model2 "url-shortener/internal/app/repository/model"
	"url-shortener/internal/config"
	"url-shortener/internal/pkg/constants"
	osfile2 "url-shortener/internal/pkg/osfile"
)

func NewURLStorage(cfg *config.AppConfig) *Storage {
	return &Storage{
		mapURL: make(map[string]string),
		BaseStorageDependency: common.BaseStorageDependency{
			Cfg: *cfg,
			Mu:  sync.RWMutex{},
		},
	}
}

func (us *Storage) Get(url string) (model2.URLData, error) {
	var URLData model2.URLData
	value, found := us.mapURL[url]
	if !found {
		return URLData, constants.ErrURLNotFound
	}

	URLData.URL = value
	return URLData, nil
}

func (us *Storage) Add(url, shortURL, userID string) (string, error) {
	value, found := us.mapURL[url]
	if found {
		return value, constants.ErrURLAlreadyExists
	}

	us.mapURL[url] = shortURL
	us.mapURL[shortURL] = url

	return "", nil
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
	data := osfile2.Event{UUID: strconv.Itoa(len(us.mapURL)), ShortURL: shortURL, OriginalURL: url}
	consumer, err := osfile2.NewConsumer(us.Cfg.FileStoragePath)
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

func (us *Storage) AddBatch(listBatches []model2.Batch, userID string) error {
	for _, batch := range listBatches {
		us.Add(batch.OriginalURL, batch.ShortURL, userID)
		err := us.Write(batch.OriginalURL, batch.ShortURL)
		if err != nil {
			return err
		}
	}

	return nil
}

func (us *Storage) GetURLsByID(userID string) ([]model2.URL, error) {
	return nil, constants.ErrUnsupportedMethod
}

func (us *Storage) BatchUpdate(event model2.DeleteEvent) error {
	return constants.ErrUnsupportedMethod
}
