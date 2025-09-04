package grpcdbhandler

import (
	"go.uber.org/zap"

	dbusecase "url-shortener/internal/app/usecase/db"
)

// New конструктор grpcdbhandler
func New(logger *zap.Logger, useCase dbusecase.IDBUseCase) *DBHandler {
	return &DBHandler{
		useCase: useCase,
		logger:  logger,
	}
}
