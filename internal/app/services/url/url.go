package url

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/storage"
	stringUtils "url-shortener/internal/app/utils/string"
)

func NewURLService(store storage.URLStorage, cfg *config.Config) *Service {
	return &Service{store: store, cfg: *cfg}
}

func (u *Service) CreateShort(url string, cfg *config.Config) string {
	var shortURL string
	var err error

	shortURL, err = u.store.Get(url)
	if err == nil {
		return shortURL
	}

	shortURL = stringUtils.CreateRandom()
	u.store.Add(url, shortURL)

	isNeedWriteToFile := cfg.DatabaseDSN == ""
	if isNeedWriteToFile {

	}

	return shortURL
}

func (u *Service) GetOriginal(shortURL string) (string, error) {
	originalURL, err := u.store.Get(shortURL)
	if err != nil {
		return "", err
	}

	//test
	//isURLEmpty := len(originalURL) == 0
	//if isURLEmpty {
	//	return "", constants.ErrURLNotFound
	//}

	return originalURL, nil
}
