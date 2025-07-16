package urlhanlder

import (
	"go.uber.org/zap"

	config "url-shortener/internal/app/config/env"
	"url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/validator"
)

type Handler struct {
	UseCase   urlusecase.IUrlUseCase
	Logger    *zap.Logger
	Validator *validator.Validator
	Config    *config.Env
}
