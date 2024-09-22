package postgres

import (
	"context"
	"database/sql"

	"github.com/felipedavid/vrcursos/src/core/model"
	"github.com/felipedavid/vrcursos/src/infrastructure/repository"
)

type PostgresCourseRepository struct {
	db *sql.DB
}

func NewPostgresCourseRepository(db *sql.DB) *PostgresCourseRepository {
	return &PostgresCourseRepository{db: db}
}

func (r PostgresCourseRepository) Save(ctx context.Context, course *model.Course) error {
	query := `INSERT INTO course (description, name) VALUES ($1, $2) RETURNING id`

	row := r.db.QueryRow(query, course.Description, course.Name)
	err := row.Scan(&course.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresCourseRepository) GetCourses(ctx context.Context) ([]*model.Course, error) {
	query := `SELECT id, description, name FROM course`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*model.Course

	for rows.Next() {
		var course model.Course
		if err := rows.Scan(&course.ID, &course.Description, &course.Name); err != nil {
			return nil, err
		}

		courses = append(courses, &course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r PostgresCourseRepository) GetCourse(ctx context.Context, id int) (*model.Course, error) {
	query := `SELECT id, description, name FROM course WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	var course model.Course
	if err := row.Scan(&course.ID, &course.Description, &course.Name); err != nil {
		return nil, err
	}

	return &course, nil
}

func (r PostgresCourseRepository) UpdateCourse(ctx context.Context, course *model.Course) error {
	query := `UPDATE course SET description = $1, name = $2 WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, course.Description, course.Name, course.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresCourseRepository) DeleteCourse(ctx context.Context, id int) error {
	query := `DELETE FROM course WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresCourseRepository) AddStudentToCourse(ctx context.Context, courseID int, studentID int) error {
	query := `INSERT INTO enrollment (course_id, student_id) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, courseID, studentID)
	if err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "unique_student_course"` {
			return repository.ErrStudentAlreadyEnrolled
		}
		return err
	}

	return nil
}

func (r PostgresCourseRepository) RemoveStudentFromCourse(ctx context.Context, courseID int, studentID int) error {
	query := `DELETE FROM enrollment WHERE course_id = $1 AND student_id = $2`

	_, err := r.db.ExecContext(ctx, query, courseID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresCourseRepository) HowManyEnrolled(ctx context.Context, courseID int) (int, error) {
	query := `SELECT COUNT(*) FROM enrollment WHERE course_id = $1`

	row := r.db.QueryRowContext(ctx, query, courseID)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
