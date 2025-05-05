package url

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"url-shortener/internal/app/middleware"
	"url-shortener/internal/app/model"
)

func (h *Handler) DeleteBatch(res http.ResponseWriter, req *http.Request) {
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

	event := model.DeleteEvent{
		ListURL: listURLs,
		UserID:  userID,
	}

	go h.Services.URL.BatchDelete(event)
	res.WriteHeader(http.StatusAccepted)
}
