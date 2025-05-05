package db

import (
	"net/http"
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/services"
)

type Method interface {
	Ping(res http.ResponseWriter, req *http.Request)
}

type Handler struct {
	Services services.Services
	common.BaseDependency
}
