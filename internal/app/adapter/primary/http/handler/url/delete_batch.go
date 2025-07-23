package urlhandler

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
)

// DeleteBatch godoc
//
//	@Summary		Deletes batch of urls from storage by user
//	@Description	DeleteBatch is used to set active flag=false for multiple url records for user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Tags			API
//	@Success		202
//	@Failure		400,401
//	@Router			/api/user/urls   [delete]
func (h *Handler) DeleteBatch(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if userID == "" && !ok {
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

	go h.useCase.DeleteBatch(event)
	res.WriteHeader(http.StatusAccepted)
}
