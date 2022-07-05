package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"reminder/config"
	"reminder/db"
	"reminder/internal/app/api/controllers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	config := config.NewConfig()
	postgres := db.NewPostgres(&config.DbConfig)
	db, err := postgres.Connect()

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	noteController := controllers.NewNoteController(db)

	router.HandleFunc("/", noteController.Index)

	http.ListenAndServe(":8080", router)
}
