package urlhandler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"url-shortener/internal/pkg/constants"
)

// GetOriginalByID godoc
// @Summary      Перенаправление на оригинальный URL
// @Description  Выполняет редирект на оригинальный URL по его сокращенному идентификатору.
//
//	Если URL был удален, возвращает статус 410 (Gone).
//
// @Tags         Перенаправления
// @Param        hash path string true "Сокращенный идентификатор URL"
// @Success      307 "Перенаправление на оригинальный URL"
// @Failure      404 {string} string "URL не найден"
// @Failure      410 {string} string "URL был удален"
// @Router       /{hash} [get]
func (h *Handler) GetOriginalByID(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	id := chi.URLParam(req, "id")
	originalURL, err := h.useCase.GetOriginalByID(ctx, id)
	if err != nil {
		h.logger.Error(errMessageFailedGetOriginalID, zap.Error(err))
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
