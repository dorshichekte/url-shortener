package urlhandler

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	errorshandler "url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/pkg/constants"
	util "url-shortener/internal/pkg/util/error_response"
)

func (h *Handler) GetStats(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	stats, err := h.useCase.GetStats(ctx)
	if err != nil {
		util.WriteErrorResponse(res, http.StatusInternalServerError, util.WrapperError[string]{CustomError: err.Error()})
		return
	}

	serviceStats := dto.ServiceStatsResponse{URLCount: stats.URLCount, UserCount: stats.UserCount}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	err = json.NewEncoder(res).Encode(serviceStats)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
	}
}
