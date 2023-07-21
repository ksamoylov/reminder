package services

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"io"
	"reminder/internal/app/models"
	"reminder/internal/app/repositories"
	"reminder/internal/app/types"
)

type NoteService struct {
	Repository *repositories.NoteRepository
}

func NewNoteService(repository *repositories.NoteRepository) *NoteService {
	return &NoteService{Repository: repository}
}

func (s *NoteService) Create(readCloser io.ReadCloser) (*models.Note, error) {
	var note models.Note

	err := json.NewDecoder(readCloser).Decode(&note)

	if err != nil {
		return nil, err
	}

	if err = validator.Validate(note); err != nil {
		return nil, err
	}

	if _, err = s.Repository.CreateOne(&note); err != nil {
		return nil, err
	}

	return &note, nil
}

func (s *NoteService) Delete(id string) error {
	var request types.IdExistingStruct

	request.Id = id

	if err := validator.Validate(request); err != nil {
		return err
	}

	if err := s.Repository.Delete(&request.Id); err != nil {
		return err
	}

	return nil
}
