package storage

import "strconv"

func NewURLStorage() *URLStorage {
	return &URLStorage{
		mapURL:      make(map[string]string),
		mapShortURL: make(map[string]string),
	}
}

func (us *URLStorage) Get(url string, urlType URLType) string {
	switch urlType {
	case DefaultURLType:
		value, found := us.mapURL[url]
		if !found {
			return ""
		}
		return value
	case ShortURLType:
		value, found := us.mapShortURL[url]
		if !found {
			return ""
		}
		return value
	default:
		return ""
	}
}

func (us *URLStorage) Add(url, shortURL string) {
	us.mapURL[url] = shortURL
	us.mapShortURL[shortURL] = url
}

func (us *URLStorage) Write(url, shortURL, fileStoragePath string) error {
	data := Event{UUID: strconv.Itoa(len(us.mapURL)), ShortURL: shortURL, OriginalURL: url}

	consumer, err := NewConsumer(fileStoragePath)
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

func (us *URLStorage) Load(fileStoragePath string) error {
	pr, err := NewProducer(fileStoragePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = pr.Close()
	}()

	_, err = pr.ReadEvent(us)
	if err != nil {
		return err
	}

	return nil
}
