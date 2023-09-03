package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"reminder/internal/app/models"
	"time"
)

const (
	TypeRefresh = "refresh"
	TypeAccess  = "access"
	Issuer      = "auth"
	Subject     = "user"
	tokenTtl    = 6
	secret      = "secret" // todo env
)

type TokenClaims struct {
	jwt.RegisteredClaims
	Type   string `json:"type"`
	UserId int    `json:"user_id"`
}

func newTokenClaims(tokenType string, userId int) *TokenClaims {
	return &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Issuer,
			Subject:   Subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTtl * time.Hour)),
		},
		Type:   tokenType,
		UserId: userId,
	}
}

type TokenPair struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

func NewTokenPair(refreshToken string, accessToken string) *TokenPair {
	return &TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
}

func CreateTokenByUser(user *models.User, tokenType string) (*string, error) {
	claims := newTokenClaims(tokenType, user.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, err
	}

	return nil, errors.New("invalid token passed")
}
