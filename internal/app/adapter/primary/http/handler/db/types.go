package dbhandler

import (
	"go.uber.org/zap"

	dbusecase "url-shortener/internal/app/usecase/db"
)

// Handler структура обработчика базы данных.
type Handler struct {
	useCase dbusecase.IDBUseCase
	logger  *zap.Logger
}
