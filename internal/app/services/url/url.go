package url

import (
	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/storage"
	stringUtils "url-shortener/internal/app/utils/string"
)

func NewURLService(store *storage.URLStorage) *Service {
	return &Service{store: *store}
}

func (u *Service) CreateShort(url string) string {
	var shortURL string

	shortURL = u.store.Get(url, storage.DefaultURLType)

	isURLEmpty := len(shortURL) == 0
	if isURLEmpty {
		shortURL = stringUtils.CreateRandom()
		u.store.Add(url, shortURL)
	}

	return shortURL
}

func (u *Service) GetOriginal(shortURL string) (string, error) {
	originalURL := u.store.Get(shortURL, storage.ShortURLType)

	isURLEmpty := len(originalURL) == 0
	if isURLEmpty {
		return "", constants.ErrURLNotFound
	}

	return originalURL, nil
}
