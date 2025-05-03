package url

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"url-shortener/internal/app/config"
	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/middleware"
	"url-shortener/internal/app/models"
	urlService "url-shortener/internal/app/services/url"
)

func NewHandler(service *urlService.Service, cfg *config.Config, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		config:  cfg,
		logger:  logger,
	}
}

func (h *Handler) handleError(res http.ResponseWriter, statusCode int) {
	res.WriteHeader(statusCode)
}

func (h *Handler) jsonDecode(req *http.Request) (models.ShortenRequest, error) {
	var request models.ShortenRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return request, err
	}

	return request, nil
}

func (h *Handler) parseRequest(req *http.Request) (string, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	if len(body) == 0 {
		return "", constants.ErrEmptyRequestBody
	}

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (h *Handler) GetURL(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	originalURL, err := h.service.GetOriginal(id)
	if err != nil {
		h.logger.Error("Failed get original URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	if originalURL.Deleted {
		res.WriteHeader(http.StatusGone)
		return
	}

	res.Header().Set("Location", originalURL.Url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) AddURL(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	originalURL, err := h.parseRequest(req)
	if err != nil {
		h.logger.Error("Failed parse request URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	shortURL, err := h.service.CreateShort(originalURL, userID)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL
	defer res.Write([]byte(fullURL))

	if err != nil {
		h.logger.Error("Failed create short URL", zap.Error(err))
		res.Header().Set("Content-Type", "text/plain")
		h.handleError(res, http.StatusConflict)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
}

func (h *Handler) AddUrlJSON(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	u, err := h.jsonDecode(req)
	if err != nil {
		h.logger.Error("Failed decode json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	shortURL, err := h.service.CreateShort(u.OriginalURL, userID)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL
	response := models.ShortenResponse{
		ShortURL: fullURL,
	}

	if err != nil {
		h.logger.Error("Failed create short URL", zap.Error(err))
		res.Header().Set("Content-Type", "application/json")
		h.handleError(res, http.StatusConflict)
		json.NewEncoder(res).Encode(response)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
}

func (h *Handler) Ping(res http.ResponseWriter, req *http.Request) {
	ps := h.config.DatabaseDSN

	db, err := sql.Open("pgx", ps)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = db.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (h *Handler) AddURLsBatch(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	var rq []models.BatchRequest
	var rs []models.BatchResponse

	body, err := io.ReadAll(req.Body)
	if err != nil {
		h.logger.Error("Failed read request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	if err := json.Unmarshal(body, &rq); err != nil {
		h.logger.Error("Failed parse request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		h.logger.Error("Empty request body")
		h.handleError(res, http.StatusBadRequest)
		return
	}

	for _, v := range rq {
		if _, err = url.ParseRequestURI(v.OriginalURL); err != nil {
			h.logger.Error("Failed parse request URL", zap.Error(err))
			h.handleError(res, http.StatusBadRequest)
			return
		}

		if v.ID == "" {
			h.logger.Error("Empty request ID")
			h.handleError(res, http.StatusBadRequest)
			return
		}
	}

	rs, err = h.service.AddBatch(rq)
	if err != nil {
		h.logger.Error("Failed add batches", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(rs)
	if err != nil {
		h.logger.Error("Failed marshal json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)

	_, err = res.Write(j)
	if err != nil {
		h.logger.Error("Failed write response", zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetURLsByID(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	listURLS, err := h.service.GetUserURLSByID(userID)
	if err != nil {
		h.logger.Error(err.Error(), zap.Error(err))
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
		h.logger.Error("Failed marshal json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(j)
	if err != nil {
		h.logger.Error("Failed write response", zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteURLsByID(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		h.logger.Error("Failed read request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	var listURLs []string
	err = json.Unmarshal(body, &listURLs)
	if err != nil {
		h.logger.Error("Failed unmarshal request body", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	go h.service.DeleteURLsByID(listURLs, userID)
	res.WriteHeader(http.StatusAccepted)
}
