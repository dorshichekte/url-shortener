package db

import (
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/services"
)

func NewDB(services services.Services, dependency common.BaseDependency) *Handler {
	return &Handler{
		Services:       services,
		BaseDependency: dependency,
	}
}
