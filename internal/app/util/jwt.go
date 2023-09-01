package util

import (
	"github.com/golang-jwt/jwt/v5"
	"reminder/internal/app/models"
)

func CreateTokenByUser(user *models.User) (*string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"name":  user.Name,
			"email": user.Email,
		})

	tokenString, err := token.SignedString([]byte(user.Password))

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
