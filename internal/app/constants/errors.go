package constants

import "errors"

var (
	ErrURLNotFound             = errors.New(urlNotFound)
	ErrEmptyRequestBody        = errors.New(emptyRequestBody)
	ErrTokenNotValid           = errors.New(tokenNotValid)
	ErrTokenHasExpired         = errors.New(tokenHasExpired)
	ErrUnexpectedSigningMethod = errors.New(unexpectedSigningMethod)
	ErrUnsupportedMethod       = errors.New(unsupportedMethod)
	ErrEmptyUserID             = errors.New(emptyUserID)
)
