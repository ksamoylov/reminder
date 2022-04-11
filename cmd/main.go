package main

import (
	"vroom/config"
	"vroom/db"
)

func main() {
	config := config.NewConfig()
	postgres := db.NewPostgres(&config.DbConfig)

	postgres.Connect()
}
