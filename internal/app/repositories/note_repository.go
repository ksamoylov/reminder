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

func (repository *NoteRepository) CreateOne(note *models.Note) *models.Note {
	var id int
	var createdAt string
	var updatedAt string

	sqlStatement := fmt.Sprintf(
		"insert into %s (name) values ($1) returning id, created_at, updated_at",
		repository.TableName,
	)

	err := repository.DB.QueryRow(sqlStatement, note.Name).Scan(&id, &createdAt, &updatedAt)

	if err != nil {
		panic(err)
	}

	note.ID = id
	note.CreatedAt = createdAt
	note.UpdatedAt = updatedAt

	return note
}

func (repository *NoteRepository) Delete(id *int) {
	sqlStatement := fmt.Sprintf(
		"delete from %s where id = %d",
		repository.TableName,
		*id,
	)

	_, err := repository.DB.Exec(sqlStatement)

	if err != nil {
		panic(err)
	}
}
