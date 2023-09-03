package types

type Success bool

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
	Success      bool   `json:"success"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewAuthTokenResponse(success bool, accessToken string, refreshToken string) *AuthTokenResponse {
	return &AuthTokenResponse{
		Success:      success,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
