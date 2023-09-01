package types

type AuthData struct {
	Email    string `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" validate:"min=8"`
}
