package urlhanlder

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) MakeFromJSON(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

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

	shortURL, err := h.UseCase.AddShorten(ctx, u.OriginalURL, userID)
	baseURL := h.Config.BaseURL
	fullURL := baseURL + "/" + shortURL
	response := dto.ShortenResponse{
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
