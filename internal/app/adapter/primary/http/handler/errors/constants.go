// Пакет errorshandler включает константы содержашие повторяющиеся текста ошибок.
package errorshandler

// Константы с текстом ошибок для обработчиков.
const (
	ErrMessageFailedParseRequestBody = "Failed to parse request body"
	ErrMessageFailedReadRequestBody  = "Failed to read request body"
	ErrMessageEmptyRequestBody       = "Empty request body"
	ErrMessageFailedParseRequestURI  = "Failed to parse request URI"
	ErrMessageFailedMarshalJSON      = "Failed to marshal json"
	ErrMessageFailedUnmarshalJSON    = "Failed to unmarshal json"
	ErrMessageFailedWriteResponse    = "Failed to write response"
	ErrMessageFailedDecodeJSON       = "Failed to decode json"
	ErrMessageFailedDecompressBody   = "Failed to decompress request body"
)
