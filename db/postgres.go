package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"vroom/config"
)

type Postgres struct {
	*config.DbConfig
}

func NewPostgres(dbConfig *config.DbConfig) *Postgres {
	return &Postgres{dbConfig}
}

func (postgres *Postgres) Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgres.DbConfig.Host,
		postgres.DbConfig.Port,
		postgres.DbConfig.User,
		postgres.DbConfig.Pass,
		postgres.DbConfig.Name,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	
	return db, nil
}
