package auth

import (
	customerror "url-shortener/internal/pkg/error"
)

var (
	errExpiredToken            = customerror.New(errMessageExpiredToken)
	errUnexpectedSigningMethod = customerror.New(errMessageUnexpectedSigningMethod)
)
