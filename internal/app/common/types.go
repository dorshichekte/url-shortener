package common

import (
	"go.uber.org/zap"
	"sync"

	"url-shortener/internal/app/config"
)

type BaseDependency struct {
	Cfg    config.Config
	Logger *zap.Logger
}

type BaseStorageDependency struct {
	Cfg config.Config
	Mu  sync.RWMutex
}
