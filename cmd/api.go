package main

import (
	"github.com/joho/godotenv"
	"reminder/config"
	"reminder/db"
	app "reminder/internal/app/server"
	"reminder/pkg/logger"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	conf := config.New()
	postgres := db.NewPostgres(&conf.DbConfig)
	dbInstance, err := postgres.Connect()

	if err != nil {
		logger.Error(err)
		return
	}

	server := app.NewServer(conf, dbInstance)
	server.Start()
}
