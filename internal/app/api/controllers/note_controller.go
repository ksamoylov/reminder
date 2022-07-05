package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"reminder/internal/app/repositories"
)

type NoteContoller struct {
	repository *repositories.NoteRepository
}

func NewNoteController(db *sql.DB) *NoteContoller {
	return &NoteContoller{
		repository: repositories.NewNoteRepository(db),
	}
}

func (controller NoteContoller) Index(w http.ResponseWriter, r *http.Request) {
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
