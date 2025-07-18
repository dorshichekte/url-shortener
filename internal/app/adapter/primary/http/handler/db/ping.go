package dbhandler

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/pkg/constants"
)

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
