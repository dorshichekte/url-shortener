package error

// TextError — строковой тип, предназначенный для описания сообщений об ошибках.
type TextError string

// CustomError — структура для представления простой кастомной ошибки с сообщением.
type CustomError struct {
	Text TextError
}

// CustomErrorWithData — структура для ошибки с сообщением и дополнительными данными.
type CustomErrorWithData struct {
	Text TextError
	Data any
}
