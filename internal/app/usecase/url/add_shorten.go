package urlusecase

import (
	"context"
	"fmt"

	customerror "url-shortener/internal/pkg/error"
	"url-shortener/internal/pkg/osfile"
	stringUtils "url-shortener/internal/pkg/util/string"
)

// ToDo поправить логику
func (u *URLUseCase) AddShorten(ctx context.Context, originalURL, userID string) (string, error) {
	shortURL := stringUtils.CreateRandom()
	url, err := u.URLRepository.AddShorten(ctx, originalURL, shortURL, userID)
	if err != nil || url != "" {
		fmt.Println(url)
		return url, customerror.New(errMessageShortURLAlreadyExists)
	}

	if u.Config.DatabaseDSN == "" {
		consumer, _ := osfile.NewConsumer(u.Config.FileStoragePath)
		_ = consumer.WriteEvent(&osfile.Event{UUID: stringUtils.CreateRandom(), OriginalURL: url, ShortURL: shortURL})
	}

	return shortURL, nil
}
