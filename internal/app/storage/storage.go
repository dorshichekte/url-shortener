package storage

import (
	"fmt"
	"go.uber.org/zap"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/osfile"
	"url-shortener/internal/app/storage/db"
	"url-shortener/internal/app/storage/memory"
)

func initDatabase(cfg *config.Config) (URLStorage, error) {
	ps, err := db.NewPostgresStorage(cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func initMemory(cfg *config.Config) (URLStorage, error) {
	st := memory.NewURLStorage(cfg)
	consumer := osfile.Consumer{}

	event, err := consumer.Load(cfg.FileStoragePath)
	if err != nil {
		return st, err
	}

	fmt.Println(event)

	return st, nil
}

func Create(cfg *config.Config, logger *zap.Logger) URLStorage {
	var store URLStorage
	var errInitDB error
	var errInitFileStorage error

	if cfg.DatabaseDSN != "" {
		store, errInitDB = initDatabase(cfg)
	}

	isFailedInitDB := errInitDB != nil || store == nil
	if isFailedInitDB {
		logger.Error("failed to connect to DB", zap.Error(errInitDB))
		store, errInitFileStorage = initMemory(cfg)

		if errInitFileStorage != nil {
			logger.Error("failed open file for memory storage", zap.Error(errInitFileStorage))
		}
	}

	return store
}
