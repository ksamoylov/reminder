package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) *StatusError {
	var response []byte

	notes, err := h.service.Repository.FindAll()

	if err != nil {
		return &StatusError{
			Err:  err,
			Code: http.StatusUnprocessableEntity,
		}
	}

	response, err = json.Marshal(notes)

	w.Write(response)

	return nil
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) *StatusError {
	var response []byte

	note, err := h.service.Create(r.Body)

	if err != nil {
		return &StatusError{
			Err:  err,
			Code: http.StatusUnprocessableEntity,
		}
	}

	response, err = json.Marshal(note)

	w.Write(response)

	return nil
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) *StatusError {
	if err := h.service.Delete(r.URL.Query().Get("id")); err != nil {
		return &StatusError{
			Err:  err,
			Code: http.StatusUnprocessableEntity,
		}
	}

	response, err := json.Marshal(CreateSuccessfulResponse())

	if err != nil {
		return &StatusError{
			Err:  err,
			Code: http.StatusUnprocessableEntity,
		}
	}

	w.Write(response)

	return nil
}
