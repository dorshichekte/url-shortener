package dbhandler

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/pkg/constants"
)

// Ping godoc
// @Summary      DB connection
// @Description  Check connection db
// @Tags         Ping
// @Success      200
// @Failure      500
// @Router       /ping [get]
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
