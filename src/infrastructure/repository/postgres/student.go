package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/felipedavid/vrcursos/src/core/model"
)

type PostgresStudentRepository struct {
	db *sql.DB
}

func NewPostgresStudentRepository(db *sql.DB) *PostgresStudentRepository {
	return &PostgresStudentRepository{db: db}
}

func (r PostgresStudentRepository) Save(ctx context.Context, student *model.Student) error {
	query := `INSERT INTO student (name) VALUES ($1) RETURNING id`

	row := r.db.QueryRow(query, student.Name)
	err := row.Scan(&student.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresStudentRepository) GetStudents(ctx context.Context, search string) ([]*model.Student, error) {
	query := `SELECT id, name FROM student`
	args := []any{}

	if search != "" {
		searchTerms := strings.Fields(search)

		query += ` WHERE `
		conditions := []string{}

		for i, term := range searchTerms {
			conditions = append(conditions, fmt.Sprintf("LOWER(UNACCENT(name)) ILIKE LOWER(UNACCENT($%d))", i+1))
			args = append(args, "%"+term+"%")
		}

		query += strings.Join(conditions, " AND ")
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*model.Student

	for rows.Next() {
		var student model.Student
		if err := rows.Scan(&student.ID, &student.Name); err != nil {
			return nil, err
		}
		students = append(students, &student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r PostgresStudentRepository) GetStudent(ctx context.Context, id int) (*model.Student, error) {
	query := `SELECT id, name FROM student WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)

	var student model.Student
	if err := row.Scan(&student.ID, &student.Name); err != nil {
		return nil, err
	}

	return &student, nil
}

func (r PostgresStudentRepository) UpdateStudent(ctx context.Context, student *model.Student) error {
	query := `UPDATE student SET name = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, student.Name, student.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresStudentRepository) DeleteStudent(ctx context.Context, id int) error {
	query := `DELETE FROM student WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresStudentRepository) EnrolledInHowManyCourses(ctx context.Context, studentID int) (int, error) {
	query := `SELECT COUNT(*) FROM enrollment WHERE student_id = $1`

	row := r.db.QueryRowContext(ctx, query, studentID)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
