package repositories

import (
	"database/sql"
	"fmt"
	"reminder/internal/app/models"
)

const UserTableName = "\"user\""

type UserRepository struct {
	DB        *sql.DB
	TableName string
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB:        db,
		TableName: UserTableName,
	}
}

func (repository *UserRepository) CreateOne(user *models.User) (*models.User, error) {
	var id int
	var createdAt string
	var updatedAt string

	sqlStatement := fmt.Sprintf(
		"insert into %s (name, password, email) values ($1, $2, $3) returning id, created_at, updated_at",
		repository.TableName,
	)

	err := repository.DB.QueryRow(sqlStatement, user.Name, user.Password, user.Email).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		return nil, err
	}

	user.ID = id
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt

	return user, nil
}

func (repository *UserRepository) CheckIfExistByEmail(email string) bool {
	var exists bool

	sqlStatement := fmt.Sprintf(
		"select exists(select * from %s where email = $1) as exists",
		repository.TableName,
	)

	err := repository.DB.QueryRow(sqlStatement, email).Scan(&exists)

	if err != nil {
		return false
	}

	return exists
}

func (repository *UserRepository) FindOneByEmail(email string) (*models.User, error) {
	var id int
	var name string
	var password string
	var createdAt string
	var updatedAt string

	sqlStatement := fmt.Sprintf(
		"select id, name, email, password, created_at, updated_at from %s where email = $1",
		repository.TableName,
	)

	err := repository.DB.QueryRow(sqlStatement, email).Scan(&id, &name, &email, &password, &createdAt, &updatedAt)

	if err != nil {
		return nil, err
	}

	var user models.User

	user.ID = id
	user.Name = name
	user.Email = email
	user.Password = password
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt

	return &user, nil
}
