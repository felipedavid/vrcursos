package postgres

import (
	"context"
	"database/sql"

	"github.com/felipedavid/vrcursos/src/core/model"
)

type PostgresStudentRepository struct {
	db *sql.DB
}

func NewPostgresStudentRepository(db *sql.DB) *PostgresStudentRepository {
	return &PostgresStudentRepository{db: db}
}

func (r PostgresStudentRepository) Save(ctx context.Context, student *model.Student) error {
	query := `INSERT INTO student (name) VALUES ($1)`
	res, err := r.db.Exec(query, student.Name)
	if err != nil {
		return err
	}

	student.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}
