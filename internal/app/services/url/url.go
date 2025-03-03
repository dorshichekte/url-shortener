package url

import (
	"errors"

	errorMessage "url-shortener/internal/app/constants"
	"url-shortener/internal/app/storage"
	stringU "url-shortener/internal/app/utils"
)

func CreateShort(url string) (string, error) {
	store := storage.GetInstance()

	if hasURL := store.Has(url, storage.DefaultUrlType); hasURL {
		return "", errors.New(errorMessage.URLAlreadyExists)
	}

	shortURL := stringU.CreateRandomString()

	store.Add(url, shortURL)

	return shortURL, nil
}

func GetOriginal(shortURL string) (string, error) {
	store := storage.GetInstance()

	if hasURL := store.Has(shortURL, storage.ShortUrlType); !hasURL {
		return "", errors.New(errorMessage.URLNotFound)
	}

	originalURL := store.Get(shortURL)

	return originalURL, nil
}
