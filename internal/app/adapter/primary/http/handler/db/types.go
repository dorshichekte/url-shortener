package dbhandler

import (
	"database/sql"

	"go.uber.org/zap"
)

// Handler структура обработчика базы данных.
type Handler struct {
	dbConnection *sql.DB
	logger       *zap.Logger
}
