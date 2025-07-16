package urlhanlder

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	config "url-shortener/internal/app/config/env"
	urlusecase "url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/constants"
	"url-shortener/internal/pkg/validator"
)

func New(logger *zap.Logger, config *config.Env, useCase urlusecase.IUrlUseCase, validator *validator.Validator) *Handler {
	return &Handler{
		UseCase:   useCase,
		Logger:    logger,
		Validator: validator,
		Config:    config,
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
		return "", constants.ErrEmptyRequestBody
	}

	_, err = url.ParseRequestURI(string(body))
	if err != nil {
		return "", err
	}

	return string(body), nil
}
