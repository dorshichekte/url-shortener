package dbhandler

import (
	"database/sql"

	"go.uber.org/zap"
)

type Handler struct {
	dbConnection *sql.DB
	logger       *zap.Logger
}
