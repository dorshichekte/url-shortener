package auth

import (
	customerror "url-shortener/internal/pkg/error"
)

var (
	errExpiredToken            = customerror.New(expiredToken)
	errInvalidToken            = customerror.New(invalidToken)
	errInitializationToken     = customerror.New(initializationToken)
	errEmptyUserID             = customerror.New(emptyUserID)
	errUnexpectedSigningMethod = customerror.New(unexpectedSigningMethod)
)
