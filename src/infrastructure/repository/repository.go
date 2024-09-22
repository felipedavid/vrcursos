package repository

import (
	"context"

	"github.com/felipedavid/vrcursos/src/core/model"
)

type IStudentRepository interface {
	Save(ctx context.Context, student *model.Student) error
	GetStudents(ctx context.Context) ([]*model.Student, error)
	GetStudent(ctx context.Context, id int) (*model.Student, error)
	UpdateStudent(ctx context.Context, student *model.Student) error
	DeleteStudent(ctx context.Context, id int) error
}

type ICourseRepository interface {
	Save(ctx context.Context, course *model.Course) error
	GetCourses(ctx context.Context) ([]*model.Course, error)
	GetCourse(ctx context.Context, id int) (*model.Course, error)
	UpdateCourse(ctx context.Context, course *model.Course) error
	DeleteCourse(ctx context.Context, id int) error
	AddStudentToCourse(ctx context.Context, courseID, studentID int) error
	RemoveStudentFromCourse(ctx context.Context, courseID, studentID int) error
	HowManyEnrolled(ctx context.Context, courseID int) (int, error)
}
