package urlusecase

import (
	"context"

	customerror "url-shortener/internal/pkg/error"
	"url-shortener/internal/pkg/osfile"
	stringUtils "url-shortener/internal/pkg/util/string"
)

// ToDo поправить логику
// AddShorten создает короткий URL для оригинального URL, если он еще не существует,
func (u *URLUseCase) AddShorten(ctx context.Context, originalURL, userID string) (string, error) {
	sU, _ := u.URLRepository.GetByOriginalURL(ctx, originalURL)
	if sU != "" {
		return sU, customerror.New(errMessageShortURLAlreadyExists)
	}

	shortURL := stringUtils.CreateRandom()
	url, addErr := u.URLRepository.AddShorten(ctx, originalURL, shortURL, userID)
	if addErr != nil {
		return url, customerror.New(errMessageShortURLAlreadyExists)
	}

	if u.Config.DatabaseDSN == "" && u.Config.FileStoragePath != "" {
		consumer, _ := osfile.NewConsumer(u.Config.FileStoragePath)
		_ = consumer.WriteEvent(&osfile.Event{UUID: stringUtils.CreateRandom(), OriginalURL: url, ShortURL: shortURL})
	}

	return shortURL, nil
}
