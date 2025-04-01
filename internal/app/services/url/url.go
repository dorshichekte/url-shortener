package url

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/storage"
	stringUtils "url-shortener/internal/app/utils/string"
)

func NewURLService(store *storage.URLStorage, cfg *config.Config) *Service {
	return &Service{store: *store, cfg: *cfg}
}

func (u *Service) CreateShort(url string, fileStoragePath string) (string, error) {
	var shortURL string
	var err error

	shortURL = u.store.Get(url, storage.DefaultURLType)

	isURLEmpty := len(shortURL) == 0
	if isURLEmpty {
		shortURL = stringUtils.CreateRandom()
		u.store.Add(url, shortURL)
		err = u.store.Write(url, shortURL, fileStoragePath)
	}

	return shortURL, err
}

func (u *Service) GetOriginal(shortURL string) (string, error) {
	originalURL := u.store.Get(shortURL, storage.ShortURLType)

	isURLEmpty := len(originalURL) == 0
	if isURLEmpty {
		return "", constants.ErrURLNotFound
	}

	return originalURL, nil
}
