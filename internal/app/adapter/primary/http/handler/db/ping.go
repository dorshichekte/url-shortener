package dbhandler

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/pkg/constants"
)

// Ping godoc
// @Summary      Проверка подключения к базе данных
// @Description  Проверяет доступность базы данных. Использует контекст с таймаутом.
// @Tags         Healthcheck
// @Produce      plain
// @Success      200  {string}  string  "База данных доступна"
// @Failure      500  {string}  string  "Ошибка подключения к базе данных"
// @Router       /ping [get]
//
// Пример:
// `GET /ping` → 200 OK (если БД доступна)
// `GET /ping` → 500 Internal Server Error (если БД недоступна)
func (h *Handler) Ping(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	if err := h.dbConnection.PingContext(ctx); err != nil {
		h.logger.Error("Failed ping db connection", zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
