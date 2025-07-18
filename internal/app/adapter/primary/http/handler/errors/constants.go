package errorshandler

const (
	ErrMessageFailedParseRequestBody = "Failed to parse request body"
	ErrMessageFailedReadRequestBody  = "Failed to read request body"
	ErrMessageEmptyRequestBody       = "Empty request body"
	ErrMessageFailedParseRequestURI  = "Failed to parse request URI"
	ErrMessageFailedMarshalJson      = "Failed to marshal json"
	ErrMessageFailedUnmarshalJson    = "Failed to unmarshal json"
	ErrMessageFailedWriteResponse    = "Failed to write response"
	ErrMessageFailedDecodeJson       = "Failed to decode json"
	ErrMessageFailedDecompressBody   = "Failed to decompress request body"
)
