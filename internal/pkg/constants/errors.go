// Пакет constants содержит глобальные константы приложения
package constants

import "errors"

// Глобальные константы ошибок.
var (
	ErrURLNotFound       = errors.New(errMessageURLNotFound)
	ErrURLAlreadyExists  = errors.New(errMessageURLAlreadyExists)
	ErrUnsupportedMethod = errors.New(errMessageUnsupportedMethod)
)
