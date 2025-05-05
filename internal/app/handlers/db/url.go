package db

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"url-shortener/internal/app/common"
	"url-shortener/internal/app/services"
)

func New(services services.Services, dependency common.BaseDependency) *Handler {
	return &Handler{
		Services:       services,
		BaseDependency: dependency,
	}
}
func (h *Handler) Ping(res http.ResponseWriter, req *http.Request) {
	ps := h.Cfg.DatabaseDSN

	db, err := sql.Open("pgx", ps)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
