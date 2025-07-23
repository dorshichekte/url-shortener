package util

import (
	"encoding/json"
	"net/http"
)

// WriteErrorResponse записывает ошибку для ответа обработчика.
func WriteErrorResponse[T ResponseTypeError](res http.ResponseWriter, statusCode int, err WrapperError[T]) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	_ = json.NewEncoder(res).Encode(&err)
}
