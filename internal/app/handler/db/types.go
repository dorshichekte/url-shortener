package db

import (
	"net/http"
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/service"
)

type Method interface {
	Ping(res http.ResponseWriter, req *http.Request)
}

type Handler struct {
	Services service.Services
	common.BaseDependency
}

func NewDB(services service.Services, dependency common.BaseDependency) *Handler {
	return &Handler{
		Services:       services,
		BaseDependency: dependency,
	}
}
