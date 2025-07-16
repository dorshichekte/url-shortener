package error

import "fmt"

func New(text string) error {
	return &CustomError{Text: TextError(text)}
}

func (e *CustomError) Error() string {
	return string(e.Text)
}

func NewWithData(text TextError, data ...any) error {
	return &CustomErrorWithData{Text: text, Data: data}
}

func (e *CustomErrorWithData) Error() string {
	return fmt.Sprintf("%s: %v", e.Text, e.Data)
}
