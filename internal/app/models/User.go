package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"min=3"`
	Email     string `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password  string `json:"password" validate:"min=8"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
