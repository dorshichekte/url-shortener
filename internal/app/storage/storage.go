package storage

import (
	"sync"
)

type URLType string

const (
	DefaultURLType URLType = "default"
	ShortURLType   URLType = "short"
)

type MapURL map[string]string

type URLStorage struct {
	mapURL      MapURL
	mapShortURL MapURL
}

type Storage interface {
	Get(shortURL string) string
	Add(url, shortURL string)
	Has(url string) (string, bool)
}

var (
	instance *URLStorage
	once     sync.Once
)

func GetInstance() *URLStorage {
	once.Do(
		func() {
			instance = &URLStorage{
				mapURL:      make(map[string]string),
				mapShortURL: make(map[string]string),
			}
		})

	return instance
}

func (us *URLStorage) Has(url string, urlType URLType) (string, bool) {
	switch urlType {
	case DefaultURLType:
		value, has := us.mapURL[url]
		return value, has
	case ShortURLType:
		value, has := us.mapShortURL[url]
		return value, has
	default:
		return "", false
	}
}

func (us *URLStorage) Get(shortURL string) string {
	return us.mapShortURL[shortURL]
}

func (us *URLStorage) Add(url, shortURL string) {
	us.mapURL[url] = shortURL
	us.mapShortURL[shortURL] = url
}
