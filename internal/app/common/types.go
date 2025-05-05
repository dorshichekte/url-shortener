package common

import (
	"sync"

	"go.uber.org/zap"

	"url-shortener/internal/app/config"
)

type BaseDependency struct {
	Cfg    config.AppConfig
	Logger *zap.Logger
}

type BaseStorageDependency struct {
	Cfg config.AppConfig
	Mu  sync.RWMutex
}
