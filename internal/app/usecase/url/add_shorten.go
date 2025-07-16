package urlusecase

import (
	"context"

	osfile2 "url-shortener/internal/pkg/osfile"
	stringUtils "url-shortener/internal/pkg/util/string"
)

func (u *URLUseCase) AddShorten(ctx context.Context, originalURL, userID string) (string, error) {
	shortURL := stringUtils.CreateRandom()
	url, err := u.URLRepository.AddShorten(ctx, originalURL, shortURL, userID)
	if err != nil {
		return url, err
	}

	if u.Config.DatabaseDSN == "" {
		consumer, _ := osfile2.NewConsumer(u.Config.FileStoragePath)
		_ = consumer.WriteEvent(&osfile2.Event{UUID: stringUtils.CreateRandom(), OriginalURL: url, ShortURL: shortURL})
	}

	return shortURL, nil
}
