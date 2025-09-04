// Пакет dbhandler включает обработчики для работы с базой данных.
package dbhandler

import (
	"go.uber.org/zap"
)

// New создаёт новый экземпляр Handler с заданными зависимостями.
func New(logger *zap.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
