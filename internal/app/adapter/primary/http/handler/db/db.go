package dbhandler

import (
	"database/sql"

	"go.uber.org/zap"
)

func New(logger *zap.Logger, dbConnection *sql.DB) *Handler {
	return &Handler{
		logger:       logger,
		dbConnection: dbConnection,
	}
}
