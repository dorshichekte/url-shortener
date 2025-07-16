package config

import (
	customerror "url-shortener/internal/pkg/error"
)

const (
	errEnvEmptyVariables customerror.TextError = "Required environment variables not found"
)
