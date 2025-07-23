package validator

import "github.com/go-playground/validator/v10"

// Validator инкапсулирует сторонний валидатор и предоставляет методы валидации и разбора ошибок.
type Validator struct {
	validator *validator.Validate
}

// ValidationError описывает структурированную ошибку валидации поля.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
