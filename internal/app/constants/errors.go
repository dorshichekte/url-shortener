package constants

import "errors"

var (
	ErrURLNotFound   = errors.New(urlNotFound)
	EmptyRequestBody = errors.New(emptyRequestBody)
)
