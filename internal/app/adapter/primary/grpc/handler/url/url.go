package grpcurlhandler

import (
	"go.uber.org/zap"

	config "url-shortener/internal/app/config/env"
	urlusecase "url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/validator"
)

// New конструктор grpcurlhandler
func New(uc urlusecase.IUrlUseCase, logger *zap.Logger, validator *validator.Validator, config *config.Env) *URLHandler {
	return &URLHandler{
		useCase:   uc,
		logger:    logger,
		validator: validator,
		config:    config,
	}
}
