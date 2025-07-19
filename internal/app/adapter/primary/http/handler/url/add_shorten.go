package urlhandler

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	errorshandler "url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) AddShorten(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if userID == "" && !ok {
		h.logger.Error(errMessageFailedGetUserIDFromContext)
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	originalURL, err := h.parseRequest(req)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedParseRequestURI, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	//ToDO переделать логику
	shortURL, err := h.useCase.AddShorten(ctx, originalURL, userID)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL
	h.logger.Info("debug test", zap.Any("url", fullURL), zap.Any("user", userID))

	if err != nil {
		h.logger.Error(errMessageFailedCreateShortURL, zap.Error(err))
		res.Header().Set("Content-Type", "text/plain")
		h.handleError(res, http.StatusConflict)

		_, jsonWriteErr := res.Write([]byte(fullURL))
		if jsonWriteErr != nil {
			h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, jsonWriteErr := res.Write([]byte(fullURL))
	if jsonWriteErr != nil {
		h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
