package url

import (
	"errors"

	errorMessage "url-shortener/internal/app/constants"
	"url-shortener/internal/app/storage"
	stringU "url-shortener/internal/app/utils"
)

func CreateShort(url string) (string, error) {
	store := storage.GetInstance()

	if hasUrl := store.Has(url, storage.DefaultUrlType); hasUrl {
		return "", errors.New(errorMessage.URLAlreadyExists)
	}

	shortURL := stringU.CreateRandomString()

	store.Add(url, shortURL)

	return shortURL, nil
}

func GetOriginal(shortUrl string) (string, error) {
	store := storage.GetInstance()

	if hasUrl := store.Has(shortUrl, storage.ShortUrlType); !hasUrl {
		return "", errors.New(errorMessage.URLNotFound)
	}

	originalUrl := store.Get(shortUrl)

	return originalUrl, nil
}
