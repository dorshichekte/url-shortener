package url

import (
	"net/http"
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/service"
)

type Method interface {
	Create(res http.ResponseWriter, req *http.Request)
	CreateBatch(res http.ResponseWriter, req *http.Request)
	CreateFromJSON(res http.ResponseWriter, req *http.Request)
	GetByID(res http.ResponseWriter, req *http.Request)
	GetAllByUser(res http.ResponseWriter, req *http.Request)
	DeleteBatch(res http.ResponseWriter, req *http.Request)
}

type Handler struct {
	Services service.Services
	common.BaseDependency
}

func NewURL(services service.Services, dependency common.BaseDependency) *Handler {
	return &Handler{
		Services:       services,
		BaseDependency: dependency,
	}
}
