package grpcurlhandler

import (
	config "url-shortener/internal/app/config/env"
	urlusecase "url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/validator"

	"go.uber.org/zap"
)

// URLHandler url хэндлер
type URLHandler struct {
	useCase   urlusecase.IUrlUseCase
	logger    *zap.Logger
	validator *validator.Validator
	config    *config.Env
}
