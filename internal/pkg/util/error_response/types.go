package util

import v "url-shortener/internal/pkg/validator"

// ResponseTypeError объявляет допустимые типы для ошибок API.
type ResponseTypeError interface {
	~string | ~[]string | ~[]v.ValidationError
}

// WrapperError универсальная обертка для ошибок API.
type WrapperError[T ResponseTypeError] struct {
	CustomError T `json:"error"`
}
