package url

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"

	"url-shortener/internal/app/middleware"
	"url-shortener/internal/app/model"
)

func (h *Handler) CreateFromJSON(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	u, err := h.jsonDecode(req)
	if err != nil {
		h.Logger.Error("Failed decode json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	shortURL, err := h.Services.URL.Shorten(u.OriginalURL, userID)
	baseURL := h.Cfg.BaseURL
	fullURL := baseURL + "/" + shortURL
	response := model.ShortenResponse{
		ShortURL: fullURL,
	}

	if err != nil {
		h.Logger.Error("Failed create short URL", zap.Error(err))
		res.Header().Set("Content-Type", "application/json")
		h.handleError(res, http.StatusConflict)
		json.NewEncoder(res).Encode(response)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
}
