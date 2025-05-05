package url

import (
	"go.uber.org/zap"
	"net/http"

	"url-shortener/internal/app/middleware"
)

func (h *Handler) Create(res http.ResponseWriter, req *http.Request) {
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

	shortURL, err := h.Services.URL.Shorten(originalURL, userID)
	baseURL := h.Cfg.BaseURL
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
