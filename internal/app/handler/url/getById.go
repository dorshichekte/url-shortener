package url

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func (h *Handler) GetByID(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	originalURL, err := h.Services.URL.GetOriginal(id)
	if err != nil {
		h.Logger.Error("Failed get original URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	if originalURL.Deleted {
		res.WriteHeader(http.StatusGone)
		return
	}

	res.Header().Set("Location", originalURL.URL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}
