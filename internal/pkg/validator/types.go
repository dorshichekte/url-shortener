package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
