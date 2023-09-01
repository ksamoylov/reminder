package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/validator.v2"
	"io"
	"reminder/internal/app/models"
	"reminder/internal/app/repositories"
	"reminder/internal/app/types"
	internalUtil "reminder/internal/app/util"
	"reminder/pkg/util"
)

type UserService struct {
	Repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{Repository: repository}
}

func (s *UserService) Create(readCloser io.ReadCloser) (*models.User, error) {
	var user models.User

	err := json.NewDecoder(readCloser).Decode(&user)

	if err != nil {
		return nil, err
	}

	if err = validator.Validate(user); err != nil {
		return nil, err
	}

	if s.Repository.CheckIfExistByEmail(user.Email) {
		return nil, errors.New(fmt.Sprintf("User %s already exists", user.Email))
	}

	user.Password, err = util.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	if _, err = s.Repository.CreateOne(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) Auth(readCloser io.ReadCloser, redis *redis.Client, ctx context.Context) (*string, error) {
	var authData types.AuthData

	err := json.NewDecoder(readCloser).Decode(&authData)

	if err != nil {
		return nil, err
	}

	if err = validator.Validate(authData); err != nil {
		return nil, err
	}

	var user *models.User

	user, err = s.checkUserPasswordByEmail(authData)

	if err != nil {
		return nil, err
	}

	var token *string
	
	token, err = internalUtil.CreateSession(user, redis, ctx)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *UserService) checkUserPasswordByEmail(authData types.AuthData) (*models.User, error) {
	user, err := s.Repository.FindOneByEmail(authData.Email)

	if err != nil {
		return nil, err
	}

	isPasswordCorrect := util.CheckPasswordHash(authData.Password, user.Password)

	if isPasswordCorrect {
		return user, nil
	}

	return nil, errors.New("invalid password received")
}
