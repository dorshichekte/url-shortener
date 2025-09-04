package grpcdbhandler

import (
	"go.uber.org/zap"

	dbusecase "url-shortener/internal/app/usecase/db"
)

// DBHandler grpc db хэндлер
type DBHandler struct {
	logger  *zap.Logger
	useCase dbusecase.IDBUseCase
}
