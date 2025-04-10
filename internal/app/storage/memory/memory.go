package memory

import (
	"fmt"
	"strconv"

	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/osfile"
)

func NewURLStorage() *Storage {
	return &Storage{
		mapURL: make(map[string]string),
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

	fmt.Println(us.mapURL)
}

func (us *Storage) Delete(url string) error {
	_, found := us.mapURL[url]
	if !found {
		return constants.ErrURLNotFound
	}

	delete(us.mapURL, url)
	return nil
}

func (us *Storage) Write(url, shortURL, fileStoragePath string) error {
	data := osfile.Event{UUID: strconv.Itoa(len(us.mapURL)), ShortURL: shortURL, OriginalURL: url}

	consumer, err := osfile.NewConsumer(fileStoragePath)
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
