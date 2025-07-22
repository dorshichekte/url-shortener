package urlhandler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	errorhandler "url-shortener/internal/app/adapter/primary/http/handler/errors"
	config "url-shortener/internal/app/config/env"
	urlusecase "url-shortener/internal/app/usecase/url"
	customerror "url-shortener/internal/pkg/error"
	"url-shortener/internal/pkg/validator"
)

func New(logger *zap.Logger, config *config.Env, useCase urlusecase.IUrlUseCase, validator *validator.Validator) *Handler {
	return &Handler{
		useCase:   useCase,
		logger:    logger,
		validator: validator,
		config:    config,
	}
}

func (h *Handler) handleError(res http.ResponseWriter, statusCode int) {
	res.WriteHeader(statusCode)
}

func (h *Handler) jsonDecode(req *http.Request) (dto.ShortenRequest, error) {
	var request dto.ShortenRequest
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

	isBodyEmpty := len(body) == 0
	if isBodyEmpty {
		return "", customerror.New(errorhandler.ErrMessageEmptyRequestBody)
	}

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}
