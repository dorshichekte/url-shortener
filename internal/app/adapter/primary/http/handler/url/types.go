package urlhandler

import (
	"go.uber.org/zap"

	config "url-shortener/internal/app/config/env"
	"url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/validator"
)

type Handler struct {
	useCase   urlusecase.IUrlUseCase
	logger    *zap.Logger
	validator *validator.Validator
	config    *config.Env
}
