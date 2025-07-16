package urlhanlder

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) GetAllByUserID(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	listURLS, err := h.UseCase.GetAllByUserID(ctx, userID)
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
