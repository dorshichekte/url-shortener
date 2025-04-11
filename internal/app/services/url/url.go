package url

import (
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/models"
	"url-shortener/internal/app/storage"
	"url-shortener/internal/app/storage/memory"
	stringUtils "url-shortener/internal/app/utils/string"
)

func NewURLService(store storage.URLStorage, cfg *config.Config) *Service {
	return &Service{store: store, cfg: *cfg}
}

func (u *Service) CreateShort(url string, cfg *config.Config) string {
	var shortURL string
	var err error

	shortURL, err = u.store.Get(url)
	if err == nil {
		return shortURL
	}

	shortURL = stringUtils.CreateRandom()
	u.store.Add(url, shortURL)

	isNeedWriteToFile := cfg.DatabaseDSN == ""
	if isNeedWriteToFile {
		m := memory.Storage{}
		_ = m.Write(url, shortURL)
	}

	return shortURL
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

	err = u.store.AddBatch(tmpListBatches)
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
