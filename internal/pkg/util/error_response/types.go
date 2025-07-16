package util

import v "url-shortener/internal/pkg/validator"

type ResponseTypeError interface {
	~string | ~[]string | ~[]v.ValidationError
}

type WrapperError[T ResponseTypeError] struct {
	CustomError T `json:"error"`
}
