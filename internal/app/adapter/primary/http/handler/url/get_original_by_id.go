package urlhanlder

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"url-shortener/internal/pkg/constants"
)

func (h *Handler) GetOriginalByID(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	id := chi.URLParam(req, "id")
	originalURL, err := h.UseCase.GetOriginalByID(ctx, id)
	if err != nil {
		h.Logger.Error("Failed get original URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	if originalURL.IsDeleted {
		res.WriteHeader(http.StatusGone)
		return
	}

	res.Header().Set("Location", originalURL.URL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
