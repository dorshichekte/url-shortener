package urlhandler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) DeleteBatch(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.logger.Error(errMessageFailedGetUserIDFromContext)
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		h.logger.Error(errorshandler.ErrMessageFailedParseRequestBody, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	var listURLs []string
	err = json.Unmarshal(body, &listURLs)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedUnmarshalJSON, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	event := entity.DeleteBatch{
		ListURL: listURLs,
		UserID:  userID,
	}

	go h.useCase.DeleteBatch(ctx, event)
	res.WriteHeader(http.StatusAccepted)
}
