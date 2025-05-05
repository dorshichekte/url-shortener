package common

import (
	"go.uber.org/zap"
	"sync"

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
