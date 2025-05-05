package url

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"url-shortener/internal/app/common"
	"url-shortener/internal/app/constants"
	"url-shortener/internal/app/middleware"
	"url-shortener/internal/app/models"
	"url-shortener/internal/app/services"
)

func NewURL(services services.Services, dependency common.BaseDependency) *Handler {
	return &Handler{
		Services:       services,
		BaseDependency: dependency,
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

func (h *Handler) GetByID(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	originalURL, err := h.Services.URL.GetOriginal(id)
	if err != nil {
		h.Logger.Error("Failed get original URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	//Тест ошибки
	//if originalURL.Deleted {
	//	res.WriteHeader(http.StatusGone)
	//	return
	//}

	res.Header().Set("Location", originalURL.URL)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func (h *Handler) Create(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	originalURL, err := h.parseRequest(req)
	if err != nil {
		h.Logger.Error("Failed parse request URL", zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	shortURL, err := h.Services.URL.Shorten(originalURL, userID)
	baseURL := h.Cfg.BaseURL
	fullURL := baseURL + "/" + shortURL
	defer res.Write([]byte(fullURL))

	if err != nil {
		h.Logger.Error("Failed create short URL", zap.Error(err))
		res.Header().Set("Content-Type", "text/plain")
		h.handleError(res, http.StatusConflict)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
}

func (h *Handler) CreateFromJSON(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	u, err := h.jsonDecode(req)
	if err != nil {
		h.Logger.Error("Failed decode json", zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	shortURL, err := h.Services.URL.Shorten(u.OriginalURL, userID)
	baseURL := h.Cfg.BaseURL
	fullURL := baseURL + "/" + shortURL
	response := models.ShortenResponse{
		ShortURL: fullURL,
	}

	if err != nil {
		h.Logger.Error("Failed create short URL", zap.Error(err))
		res.Header().Set("Content-Type", "application/json")
		h.handleError(res, http.StatusConflict)
		json.NewEncoder(res).Encode(response)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(response)
}

func (h *Handler) CreateBatch(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey()).(string)
	if !ok || userID == "" {
		h.Logger.Error("Failed get userID from context")
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	var rq []models.BatchRequest
	var rs []models.BatchResponse

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

	event := models.DeleteEvent{
		ListURL: listURLs,
		UserID:  userID,
	}

	go h.Services.URL.BatchDelete(context.Background(), event)
	res.WriteHeader(http.StatusAccepted)
}
