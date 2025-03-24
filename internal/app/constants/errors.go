package constants

import "errors"

var (
	ErrURLNotFound      = errors.New(urlNotFound)
	ErrEmptyRequestBody = errors.New(emptyRequestBody)
)
