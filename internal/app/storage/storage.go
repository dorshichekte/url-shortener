package storage

import "sync"

type UrlType string

const (
	DefaultUrlType UrlType = "default"
	ShortUrlType   UrlType = "short"
)

type MapUrl map[string]string

type URLStorage struct {
	mapUrl      MapUrl
	mapShortUrl MapUrl
}

type Storage interface {
	Get(shortUrl string) string
	Add(url, shortUrl string)
	Has(url string) bool
}

var (
	instance *URLStorage
	once     sync.Once
)

func GetInstance() *URLStorage {
	once.Do(
		func() {
			instance = &URLStorage{
				mapUrl:      make(map[string]string),
				mapShortUrl: make(map[string]string),
			}
		})

	return instance
}

func (us *URLStorage) Has(url string, urlType UrlType) bool {
	switch urlType {
	case DefaultUrlType:
		_, has := us.mapUrl[url]
		return has
	case ShortUrlType:
		_, has := us.mapShortUrl[url]
		return has
	default:
		return false
	}
}

func (us *URLStorage) Get(shortUrl string) string {
	return us.mapShortUrl[shortUrl]
}

func (us *URLStorage) Add(url, shortUrl string) {
	us.mapUrl[url] = shortUrl
	us.mapShortUrl[shortUrl] = url
}
