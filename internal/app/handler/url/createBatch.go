package url

import (
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"

	"url-shortener/internal/app/middleware"
	"url-shortener/internal/app/model"
)

func (h *Handler) CreateBatch(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	var rq []model.BatchRequest
	var rs []model.BatchResponse

	body, err := io.ReadAll(req.Body)
	if err != nil {
		h.Logger.Error("Failed read request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	if err := json.Unmarshal(body, &rq); err != nil {
		h.Logger.Error("Failed parse request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		h.Logger.Error("Empty request body")
		h.handleError(res, http.StatusBadRequest)
		return
	}

	for _, v := range rq {
		if _, err = url.ParseRequestURI(v.OriginalURL); err != nil {
			h.Logger.Error("Failed parse request URL", zap.Error(err))
			h.handleError(res, http.StatusBadRequest)
			return
		}

		if v.ID == "" {
			h.Logger.Error("Empty request ID")
			h.handleError(res, http.StatusBadRequest)
			return
		}
	}

	rs, err = h.Services.URL.BatchShorten(rq)
	if err != nil {
		h.Logger.Error("Failed add batches", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(rs)
	if err != nil {
		h.Logger.Error("Failed marshal json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)

	_, err = res.Write(j)
	if err != nil {
		h.Logger.Error("Failed write response", zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
