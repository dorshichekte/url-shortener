package httpadapter

import server "url-shortener/internal/pkg/server/http"

// HTTPAdapter сервер httpadapter
type HTTPAdapter struct {
	server *server.Http
}
