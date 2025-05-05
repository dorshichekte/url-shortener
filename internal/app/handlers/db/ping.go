package db

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"go.uber.org/zap"

	"url-shortener/internal/app/constants"
)

func (h *Handler) Ping(res http.ResponseWriter, req *http.Request) {
	ps := h.Cfg.DatabaseDSN

	db, err := sql.Open("pgx", ps)
	if err != nil {
		h.Logger.Error(constants.ErrFailedConnectionDB.Error(), zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		h.Logger.Error(constants.ErrPingTimeout.Error(), zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
