package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reminder/internal/app/repositories"
	"reminder/internal/app/services"
	"strconv"
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

func (controller *NoteContoller) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
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

func (controller *NoteContoller) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

func (controller *NoteContoller) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)

		// todo validation
		w.Write([]byte("Bad request"))

		return
	}

	formattedId, err := strconv.Atoi(id)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

		return
	}

	controller.repository.Delete(&formattedId)

	response := map[string]string{"success": "true", "message": fmt.Sprintf("Note %d deleted", formattedId)}
	responseJson, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
