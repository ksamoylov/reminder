package handlers

import (
	"database/sql"
	"gopkg.in/validator.v2"
	"net/http"
	"reminder/config"
	"reminder/internal/app/repositories"
	"reminder/internal/app/services"
)

type StatusError struct {
	Err  error
	Code int
}

type Handler struct {
	config    *config.Config
	validator *validator.Validator
	service   *services.NoteService
}

func NewHandler(config *config.Config, db *sql.DB) *Handler {
	repository := repositories.NewNoteRepository(db)
	service := services.NewNoteService(repository)

	return &Handler{config: config, service: service}
}

type HandlerFn func(w http.ResponseWriter, r *http.Request) *StatusError

func (fn HandlerFn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setDefaultResponseHeaders(w)
	statusError := fn(w, r)

	if statusError != nil {
		handleError(w, statusError)
	}
}

func setDefaultResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
