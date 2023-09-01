package handlers

import (
	"encoding/json"
	"net/http"
	"reminder/internal/app/types"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) *types.StatusError {
	var response []byte

	notes, err := h.deps.NoteService.Repository.FindAll()

	if err != nil {
		return types.NewStatusError(err, http.StatusUnprocessableEntity)
	}

	response, err = json.Marshal(notes)

	w.Write(response)

	return nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *types.StatusError {
	var response []byte

	note, err := h.deps.NoteService.Create(r.Body)

	if err != nil {
		return types.NewStatusError(err, http.StatusUnprocessableEntity)
	}

	response, err = json.Marshal(note)

	w.Write(response)

	return nil
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) *types.StatusError {
	if err := h.deps.NoteService.Delete(r.URL.Query().Get("id")); err != nil {
		return types.NewStatusError(err, http.StatusUnprocessableEntity)
	}

	response, err := json.Marshal(CreateSuccessfulResponse("Note successfully deleted"))

	if err != nil {
		return types.NewStatusError(err, http.StatusUnprocessableEntity)
	}

	w.Write(response)

	return nil
}
