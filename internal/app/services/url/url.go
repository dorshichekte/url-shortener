package url

import (
	"errors"
	"fmt"

	errorMessage "url-shortener/internal/app/constants"
	"url-shortener/internal/app/storage"
	stringU "url-shortener/internal/app/utils"
)

func CreateShort(url string) string {
	store := storage.GetInstance()
	var shortURL string

	shortURL, has := store.Has(url, storage.DefaultURLType)
	if has {
		return shortURL
	}

	shortURL = stringU.CreateRandomString()

	store.Add(url, shortURL)

	return shortURL
}

func GetOriginal(shortURL string) (string, error) {
	store := storage.GetInstance()

	if _, hasURL := store.Has(shortURL, storage.ShortURLType); !hasURL {
		return "", errors.New(errorMessage.URLNotFound)
	}

	originalURL := store.Get(shortURL)
	return originalURL, nil
}
