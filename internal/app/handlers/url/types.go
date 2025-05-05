package url

import (
	"net/http"
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/services"
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
	Services services.Services
	common.BaseDependency
}
