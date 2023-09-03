package util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"reminder/internal/app/models"
	"time"
)

const sessionKeyTemplate = "session-%d"

func createSessionKey(userId int) string {
	return fmt.Sprintf(sessionKeyTemplate, userId)
}

func CreateSession(user *models.User, redis *redis.Client, ctx context.Context) (*string, *string, error) {
	refreshToken, err := CreateTokenByUser(user, TypeRefresh)

	if err != nil {
		return nil, nil, err
	}

	var accessToken *string

	accessToken, err = CreateTokenByUser(user, TypeAccess)

	if err != nil {
		return nil, nil, err
	}

	var cachedTokenPair []byte

	cachedTokenPair, err = json.Marshal(NewTokenPair(*refreshToken, *accessToken))
	duration := tokenTtl * time.Hour

	err = redis.
		Set(ctx, createSessionKey(user.ID), string(cachedTokenPair), duration).
		Err()

	if err != nil {
		return nil, nil, err
	}

	return accessToken, refreshToken, nil
}

func CheckSessionOnExists(claims *TokenClaims, redis *redis.Client, ctx context.Context) bool {
	_, err := redis.Get(ctx, createSessionKey(claims.UserId)).Result()

	if err != nil {
		return false
	}

	return true
}
