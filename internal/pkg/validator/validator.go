// Пакет validator инициализирует валидатор, и предоставляет методы для работы с ним.
package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"

	customerror "url-shortener/internal/pkg/error"
)

// New создает и возвращает новый экземпляр Validator с инициализированным валидатором.
func New() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// ValidateStruct запускает валидацию структуры на основе тегов `validate`.
func (v *Validator) ValidateStruct(s any) error {
	return v.validator.Struct(s)
}

// ParseValidationErrors преобразует ошибки, возвращенные валидатором, в слайс ValidationError.
func (v *Validator) ParseValidationErrors(err error) ([]ValidationError, error) {
	var listErrors []ValidationError

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			ce := v.CreateError(e)
			listErrors = append(listErrors, ce)
		}

		return listErrors, nil
	}

	return nil, customerror.New(errMessageFailedParseValidationErrors)
}

// CreateError создает кастомную ValidationError на основе типа ошибки из валидатора.
func (v *Validator) CreateError(e validator.FieldError) ValidationError {
	field := strings.ToLower(e.Field())

	switch e.Tag() {
	case "required":
		return ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s is required", field),
		}
	case "min":
		return ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be at least %s characters long.", field, e.Param()),
		}
	case "max":
		return ValidationError{
			Field:   field,
			Message: fmt.Sprintf("%s must be at most %s characters long.", field, e.Param()),
		}
	default:
		return ValidationError{
			Field:   field,
			Message: e.Error(),
		}
	}
}
