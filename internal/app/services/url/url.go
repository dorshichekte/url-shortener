package url

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/models"
	"url-shortener/internal/app/osfile"
	"url-shortener/internal/app/storage"
	stringUtils "url-shortener/internal/app/utils/string"
)

func NewURLService(store storage.URLStorage, cfg *config.Config) *Service {
	return &Service{store: store, cfg: *cfg}
}

func (u *Service) CreateShort(url string) (string, error) {
	shortURL := stringUtils.CreateRandom()
	err := u.store.Add(url, shortURL, "")
	if err != nil {
		return url, err
	}

	if u.cfg.DatabaseDSN == "" {
		consumer, _ := osfile.NewConsumer(u.cfg.FileStoragePath)
		_ = consumer.WriteEvent(&osfile.Event{UUID: stringUtils.CreateRandom(), OriginalURL: url, ShortURL: shortURL})
	}

	return shortURL, nil
}

func (u *Service) GetOriginal(shortURL string) (string, error) {
	originalURL, err := u.store.Get(shortURL)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}

func (u *Service) AddBatch(listBatches []models.BatchRequest) ([]models.BatchResponse, error) {
	var err error

	tmpListBatches := make([]models.Batch, 0, len(listBatches))

	for _, batch := range listBatches {
		shortURL := stringUtils.CreateRandom()

		tmpListBatches = append(tmpListBatches, models.Batch{
			OriginalURL: batch.OriginalURL,
			ID:          batch.ID,
			ShortURL:    shortURL,
		})
	}

	err = u.store.AddBatch(tmpListBatches, "")
	if err != nil {
		return nil, err
	}

	listResponseBatches := make([]models.BatchResponse, 0, len(listBatches))
	for _, batch := range tmpListBatches {
		listResponseBatches = append(listResponseBatches, models.BatchResponse{
			ID:       batch.ID,
			ShortURL: u.cfg.BaseURL + "/" + batch.ShortURL,
		})
	}

	return listResponseBatches, nil
}

func (u *Service) GetUserURLSByID(userID string) ([]models.URL, error) {
	return u.store.GetUsersURLsByID(userID)
}
