package models

type Note struct {
	ID        int    `json:"id"`
	Name      string `json:"name" validate:"min=3"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
