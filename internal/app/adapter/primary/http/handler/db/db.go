// Пакет dbhandler включает обработчики для работы с базой данных.
package dbhandler

import (
	"database/sql"

	"go.uber.org/zap"
)

// New создаёт новый экземпляр Handler с заданными зависимостями.
func New(logger *zap.Logger, dbConnection *sql.DB) *Handler {
	return &Handler{
		logger:       logger,
		dbConnection: dbConnection,
	}
}
