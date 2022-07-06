package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reminder/internal/app"
	"reminder/internal/app/repositories"
	"reminder/internal/app/services"
)

type NoteContoller struct {
	repository *repositories.NoteRepository
	service    *services.NoteService
}

func NewNoteController(db *sql.DB) *NoteContoller {
	repository := repositories.NewNoteRepository(db)
	service := services.NewNoteService(repository)

	return &NoteContoller{
		repository: repository,
		service:    service,
	}
}

func (controller NoteContoller) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != app.POST {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	notes, err := controller.repository.FindAll()

	if err != nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	response, err := json.Marshal(notes)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (controller NoteContoller) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != app.POST {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		// todo validation
		w.Write([]byte("Bad request"))

		return
	}

	note := controller.service.Create(body)

	jsonNote, err := json.Marshal(note)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonNote)
}
