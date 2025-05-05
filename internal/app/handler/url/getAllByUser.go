package url

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"url-shortener/internal/app/middleware"
)

func (h *Handler) GetAllByUser(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	listURLS, err := h.Services.URL.GetByUserID(userID)
	if err != nil {
		h.Logger.Error(err.Error(), zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	isListURLSEmpty := len(listURLS) == 0
	if isListURLSEmpty {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	j, err := json.Marshal(listURLS)
	if err != nil {
		h.Logger.Error("Failed marshal json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(j)
	if err != nil {
		h.Logger.Error("Failed write response", zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
	}
}
