package types

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewResponse(success bool, message string) *Response {
	return &Response{
		Success: success,
		Message: message,
	}
}
