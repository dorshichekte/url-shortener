package urlusecase

import (
	"context"

	"url-shortener/internal/pkg/osfile"
	stringUtils "url-shortener/internal/pkg/util/string"
)

func (u *URLUseCase) AddShorten(ctx context.Context, originalURL, userID string) (string, error) {
	shortURL := stringUtils.CreateRandom()
	url, err := u.URLRepository.AddShorten(ctx, originalURL, shortURL, userID)
	if err != nil {
		return url, err
	}

	if u.Config.DatabaseDSN == "" {
		consumer, _ := osfile.NewConsumer(u.Config.FileStoragePath)
		_ = consumer.WriteEvent(&osfile.Event{UUID: stringUtils.CreateRandom(), OriginalURL: url, ShortURL: shortURL})
	}

	return shortURL, nil
}
