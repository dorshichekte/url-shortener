package url

import (
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/model"
	"url-shortener/internal/app/osfile"
	"url-shortener/internal/app/storage"
	stringUtils "url-shortener/internal/app/util/string"
)

func NewService(store storage.URLStorage, baseDependency common.BaseDependency) Methods {
	return &Service{
		Store:          store,
		BaseDependency: baseDependency,
	}
}

func (u *Service) Shorten(url, userID string) (string, error) {
	shortURL := stringUtils.CreateRandom()
	url, err := u.Store.Add(url, shortURL, userID)
	if err != nil {
		return url, err
	}

	if u.Cfg.DatabaseDSN == "" {
		consumer, _ := osfile.NewConsumer(u.Cfg.FileStoragePath)
		_ = consumer.WriteEvent(&osfile.Event{UUID: stringUtils.CreateRandom(), OriginalURL: url, ShortURL: shortURL})
	}

	return shortURL, nil
}

func (u *Service) GetOriginal(shortURL string) (model.URLData, error) {
	URLData, err := u.Store.Get(shortURL)
	if err != nil {
		return URLData, err
	}

	return URLData, nil
}

func (u *Service) BatchShorten(listBatches []model.BatchRequest) ([]model.BatchResponse, error) {
	var err error

	tmpListBatches := make([]model.Batch, 0, len(listBatches))

	for _, batch := range listBatches {
		shortURL := stringUtils.CreateRandom()

		tmpListBatches = append(tmpListBatches, model.Batch{
			OriginalURL: batch.OriginalURL,
			ID:          batch.ID,
			ShortURL:    shortURL,
		})
	}

	err = u.Store.AddBatch(tmpListBatches, "")
	if err != nil {
		return nil, err
	}

	listResponseBatches := make([]model.BatchResponse, 0, len(listBatches))
	for _, batch := range tmpListBatches {
		listResponseBatches = append(listResponseBatches, model.BatchResponse{
			ID:       batch.ID,
			ShortURL: u.Cfg.BaseURL + "/" + batch.ShortURL,
		})
	}

	return listResponseBatches, nil
}

func (u *Service) GetByUserID(userID string) ([]model.URL, error) {
	return u.Store.GetURLsByID(userID)
}

func (u *Service) BatchDelete(event model.DeleteEvent) error {
	return u.Store.BatchUpdate(event)
}
