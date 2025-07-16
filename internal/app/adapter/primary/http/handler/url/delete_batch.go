package urlhanlder

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) DeleteBatch(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		h.Logger.Error("Failed read request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	var listURLs []string
	err = json.Unmarshal(body, &listURLs)
	if err != nil {
		h.Logger.Error("Failed unmarshal request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	event := entity.DeleteBatch{
		ListURL: listURLs,
		UserID:  userID,
	}

	go h.UseCase.DeleteBatch(ctx, event)
	res.WriteHeader(http.StatusAccepted)
}
