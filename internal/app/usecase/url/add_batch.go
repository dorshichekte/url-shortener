package urlusecase

import (
	"context"

	entity "url-shortener/internal/app/domain/entity/url"
	stringUtils "url-shortener/internal/pkg/util/string"
)

func (u *URLUseCase) AddBatch(ctx context.Context, batches []entity.Batch, userID string) ([]entity.Batch, error) {
	tmpListBatches := make([]entity.Batch, 0, len(batches))

	for _, batch := range batches {
		shortURL := stringUtils.CreateRandom()

		tmpListBatches = append(tmpListBatches, entity.Batch{
			OriginalURL: batch.OriginalURL,
			ID:          batch.ID,
			ShortURL:    shortURL,
		})
	}

	err := u.URLRepository.AddBatch(ctx, tmpListBatches, userID)
	if err != nil {
		return nil, err
	}

	listResponseBatches := make([]entity.Batch, 0, len(batches))
	for _, batch := range tmpListBatches {
		listResponseBatches = append(listResponseBatches, entity.Batch{
			ID:       batch.ID,
			ShortURL: u.Config.BaseURL + "/" + batch.ShortURL,
		})
	}

	return listResponseBatches, nil
}
