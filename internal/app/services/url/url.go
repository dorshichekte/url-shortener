package url

import (
	"log"
	"strconv"
	"time"

	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/osfile"
	"url-shortener/internal/app/storage"
	stringUtils "url-shortener/internal/app/utils/string"
)

func NewURLService(store *storage.URLStorage) *Service {
	return &Service{store: *store}
}

func (u *Service) CreateShort(url string, fileStoragePath string) string {
	var shortURL string

	shortURL = u.store.Get(url, storage.DefaultURLType)

	isURLEmpty := len(shortURL) == 0
	if isURLEmpty {
		shortURL = stringUtils.CreateRandom()
		u.store.Add(url, shortURL)

		producer, err := osfile.NewProducer(fileStoragePath)
		if err != nil {
			log.Printf("Error init producer: %v\n", err)
		} else {
			currentTime := time.Now()
			intFromTime := currentTime.Unix()
			event := osfile.Event{
				UUID:        strconv.Itoa(int(intFromTime)),
				ShortURL:    shortURL,
				OriginalURL: url,
			}

			err = producer.WriteEvent(&event)
			if err != nil {
				log.Printf("Error write in file: %v\n", err)
			}

			log.Println("Success write in file", fileStoragePath)
			defer producer.Close()
		}
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
