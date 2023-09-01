package util

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"reminder/internal/app/models"
	"time"
)

const SessionExpirationHours = 6

func CreateSession(user *models.User, redis *redis.Client, ctx context.Context) (*string, error) {
	token, err := CreateTokenByUser(user)

	if err != nil {
		return nil, err
	}

	duration := SessionExpirationHours * time.Hour
	expiresIn := time.Now().Add(duration).Unix()

	err = redis.
		Set(ctx, *token, fmt.Sprintf("expires_in: %d", expiresIn), duration).
		Err()

	if err != nil {
		return nil, err
	}

	return token, nil
}
