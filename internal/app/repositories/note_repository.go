package repositories

import (
	"database/sql"
	"fmt"
	"reminder/internal/app/models"
)

var TableName = "note"

type NoteRepository struct {
	DB        *sql.DB
	TableName string
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{
		DB:        db,
		TableName: TableName,
	}
}

func (repository *NoteRepository) FindAll() ([]models.Note, error) {
	sqlStatement := fmt.Sprintf(`select id, name, created_at, updated_at from %s`, repository.TableName)

	rows, err := repository.DB.Query(sqlStatement)

	defer rows.Close()

	var notes []models.Note

	for rows.Next() {
		var note models.Note

		if err := rows.Scan(
			&note.ID,
			&note.Name,
			&note.CreatedAt,
			&note.UpdatedAt,
		); err != nil {
			return notes, err
		}

		notes = append(notes, note)
	}

	if err = rows.Err(); err != nil {
		return notes, err
	}

	return notes, nil
}
