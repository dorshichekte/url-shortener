package error

type TextError string

type CustomError struct {
	Text TextError
}

type CustomErrorWithData struct {
	Text TextError
	Data any
}
