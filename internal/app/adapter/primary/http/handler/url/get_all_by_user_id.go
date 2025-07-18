package urlhandler

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) GetAllByUserID(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
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
		res.WriteHeader(http.StatusNoContent)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	err = json.NewEncoder(res).Encode(listURLS)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
	}
}
