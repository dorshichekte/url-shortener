package urlhandler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	"url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) AddBatch(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if userID == "" && !ok {
		h.logger.Error(errMessageFailedGetUserIDFromContext)
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedReadRequestBody, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	var batchesRequest []dto.BatchRequest

	if err := json.Unmarshal(body, &batchesRequest); err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedUnmarshalJSON, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		h.logger.Error(errorshandler.ErrMessageEmptyRequestBody)
		h.handleError(res, http.StatusBadRequest)
		return
	}

	for _, batchRequest := range batchesRequest {
		if _, err = url.ParseRequestURI(batchRequest.OriginalURL); err != nil {
			h.logger.Error(errorshandler.ErrMessageFailedParseRequestURI, zap.Error(err))
			h.handleError(res, http.StatusBadRequest)
			return
		}

		if batchRequest.ID == "" {
			h.logger.Error(errMessageEmptyBatchRequestID)
			h.handleError(res, http.StatusBadRequest)
			return
		}
	}

	//ToDo добавить методы маппинга дто, переработать ошибки, добавить метод
	var batchesEntity = make([]entity.Batch, 0, len(batchesRequest))
	for _, batchRequest := range batchesRequest {
		batch := entity.Batch{OriginalURL: batchRequest.OriginalURL, ID: batchRequest.ID}
		batchesEntity = append(batchesEntity, batch)
	}

	batches, err := h.useCase.AddBatch(ctx, batchesEntity, userID)
	if err != nil {
		h.logger.Error(errMessageFailedAddBatches, zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	var batchResponse []dto.BatchResponse
	for _, batch := range batches {
		batchResponse = append(batchResponse, dto.BatchResponse{ID: batch.ID, ShortURL: batch.ShortURL})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(res).Encode(batchResponse)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
