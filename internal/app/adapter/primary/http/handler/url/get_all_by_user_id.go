package urlhandler

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	"url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

// GetAllByUserID godoc
//
//	@Summary		Retrieves all urls saved by user
//	@Description	GetAllByUserID retrieves all active urls saved by user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Tags			API
//	@Success		200 {object}	dto.URLRequest
//	@Success		204
//	@Failure		401,500
//	@Router			/api/user/urls   [get]
func (h *Handler) GetAllByUserID(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if userID == "" && !ok {
		h.logger.Error(errMessageFailedGetUserIDFromContext)
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	listURLS, err := h.useCase.GetAllByUserID(ctx, userID)
	if err != nil {
		h.logger.Error(err.Error(), zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	isListURLSEmpty := len(listURLS) == 0
	if isListURLSEmpty {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNoContent)
		return
	}

	var urlData []dto.URLRequest
	for _, url := range listURLS {
		urlData = append(urlData, dto.URLRequest{ShortURL: url.ShortURL, OriginalURL: url.OriginalURL})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	err = json.NewEncoder(res).Encode(urlData)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
	}
}
