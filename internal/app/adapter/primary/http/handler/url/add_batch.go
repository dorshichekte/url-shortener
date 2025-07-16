package urlhanlder

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/pkg/constants"
)

func (h *Handler) AddBatch(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		h.Logger.Error("Failed read request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	var batchesRequest []dto.BatchRequest

	if err := json.Unmarshal(body, &batchesRequest); err != nil {
		h.Logger.Error("Failed parse request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		h.Logger.Error("Empty request body")
		h.handleError(res, http.StatusBadRequest)
		return
	}

	for _, batchRequest := range batchesRequest {
		if _, err = url.ParseRequestURI(batchRequest.OriginalURL); err != nil {
			h.Logger.Error("Failed parse request URLRequest", zap.Error(err))
			h.handleError(res, http.StatusBadRequest)
			return
		}

		if batchRequest.ID == "" {
			h.Logger.Error("Empty request ID")
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

	batches, err := h.UseCase.AddBatch(ctx, batchesEntity, userID)
	if err != nil {
		h.Logger.Error("Failed add batches", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	var batchResponse []dto.BatchResponse
	for _, batch := range batches {
		batchResponse = append(batchResponse, dto.BatchResponse{ID: batch.ID, ShortURL: batch.ShortURL})
	}

	marshaledBatches, err := json.Marshal(batchResponse)
	if err != nil {
		h.Logger.Error("Failed marshal json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)

	_, err = res.Write(marshaledBatches)
	if err != nil {
		h.Logger.Error("Failed write response", zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
