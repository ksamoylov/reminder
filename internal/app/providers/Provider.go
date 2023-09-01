package providers

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gopkg.in/validator.v2"
	"reminder/internal/app/repositories"
	"reminder/internal/app/services"
)

type Repos struct {
	UserRepo *repositories.UserRepository
	NoteRepo *repositories.NoteRepository
}

func NewRepos(db *sql.DB) *Repos {
	return &Repos{
		UserRepo: repositories.NewUserRepository(db),
		NoteRepo: repositories.NewNoteRepository(db),
	}
}

type Services struct {
	UserService *services.UserService
	NoteService *services.NoteService
}

func NewServices(repos *Repos) *Services {
	return &Services{
		UserService: services.NewUserService(repos.UserRepo),
		NoteService: services.NewNoteService(repos.NoteRepo),
	}
}

type Deps struct {
	*Repos
	*Services
	*validator.Validator
	Redis *redis.Client
}

func NewDeps(db *sql.DB, redis *redis.Client) *Deps {
	repos := NewRepos(db)

	return &Deps{
		Repos:     repos,
		Services:  NewServices(repos),
		Validator: validator.NewValidator(),
		Redis:     redis,
	}
}
