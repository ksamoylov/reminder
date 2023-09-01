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

type StatusError struct {
	Err  error
	Code int
}

func NewStatusError(err error, code int) *StatusError {
	return &StatusError{
		Err:  err,
		Code: code,
	}
}

type AuthTokenResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func NewAuthTokenResponse(success bool, token string) *AuthTokenResponse {
	return &AuthTokenResponse{
		Success: success,
		Token:   token,
	}
}
