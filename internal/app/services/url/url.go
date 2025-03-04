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

	originalURL, hasURL := store.Has(shortURL, storage.ShortURLType)
	if !hasURL {
		fmt.Println("short url does not exist")
		return "", errors.New(errorMessage.URLNotFound)
	}

	return originalURL, nil
}
