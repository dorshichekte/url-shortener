package constants

import "errors"

var (
	ErrURLNotFound       = errors.New(errMessageURLNotFound)
	ErrURLAlreadyExists  = errors.New(errMessageURLAlreadyExists)
	ErrUnsupportedMethod = errors.New(errMessageUnsupportedMethod)
)
