// Пакет error предоставляет методы для создания кастомных ошибок.
package error

import "fmt"

// New создает новую ошибку с текстом, обернутым в тип TextError.
func New(text string) error {
	return &CustomError{Text: TextError(text)}
}

// Error возвращает строковое представление ошибки CustomError.
func (e *CustomError) Error() string {
	return string(e.Text)
}

// NewWithData создает ошибку с текстом и дополнительными данными.
func NewWithData(text TextError, data ...any) error {
	return &CustomErrorWithData{Text: text, Data: data}
}

// Error возвращает строковое представление ошибки CustomErrorWithData
func (e *CustomErrorWithData) Error() string {
	return fmt.Sprintf("%s: %v", e.Text, e.Data)
}
