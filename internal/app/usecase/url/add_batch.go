package urlusecase

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
	stringUtils "url-shortener/internal/pkg/util/string"
)

func (u *URLUseCase) AddBatch(ctx context.Context, batches []entity.Batch, userID string) ([]entity.Batch, error) {
	temporaryBatches := make([]entity.Batch, 0, len(batches))

	for _, batch := range batches {
		shortURL := stringUtils.CreateRandom()

		temporaryBatches = append(temporaryBatches, entity.Batch{
			OriginalURL: batch.OriginalURL,
			ID:          batch.ID,
			ShortURL:    shortURL,
		})
	}

	err := u.URLRepository.AddBatch(ctx, temporaryBatches, userID)
	if err != nil {
		return nil, err
	}

	listResponseBatches := make([]entity.Batch, 0, len(batches))
	for _, batch := range temporaryBatches {
		listResponseBatches = append(listResponseBatches, entity.Batch{
			ID:       batch.ID,
			ShortURL: u.Config.BaseURL + "/" + batch.ShortURL,
		})
	}

	return listResponseBatches, nil
}
