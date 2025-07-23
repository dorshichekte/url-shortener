// Пакет logger инициализирует логгер.
package logger

import (
	"go.uber.org/zap"
)

// New создает экземпляр логгера.
func New() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
