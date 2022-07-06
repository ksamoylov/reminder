package services

import (
	"encoding/json"
	"reminder/internal/app/models"
	"reminder/internal/app/repositories"
)

type NoteService struct {
	repository *repositories.NoteRepository
}

func NewNoteService(repository *repositories.NoteRepository) *NoteService {
	return &NoteService{repository: repository}
}

func (service *NoteService) Create(data []byte) models.Note {
	var note models.Note

	json.Unmarshal(data, &note)

	service.repository.CreateOne(&note)

	return note
}
