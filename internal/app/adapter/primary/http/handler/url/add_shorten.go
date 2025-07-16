package urlhanlder

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) AddShorten(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	originalURL, err := h.parseRequest(req)
	if err != nil {
		h.Logger.Error("Failed parse request URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	shortURL, err := h.UseCase.AddShorten(ctx, originalURL, userID)
	baseURL := h.Config.BaseURL
	fullURL := baseURL + "/" + shortURL
	if err != nil {
		h.Logger.Error("Failed create short URL", zap.Error(err))
		res.Header().Set("Content-Type", "text/plain")
		h.handleError(res, http.StatusConflict)
		_, _ = res.Write([]byte(fullURL))
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, _ = res.Write([]byte(fullURL))
}
