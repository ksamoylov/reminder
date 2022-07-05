package main

import (
	"reminder/config"
	"reminder/db"
)

func main() {
	config := config.NewConfig()
	postgres := db.NewPostgres(&config.DbConfig)

	postgres.Connect()
}
