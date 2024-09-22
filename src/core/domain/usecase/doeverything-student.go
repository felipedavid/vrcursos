package usecase

import (
	"context"

	"github.com/felipedavid/vrcursos/src/core/domain"
	"github.com/felipedavid/vrcursos/src/core/model"
	"github.com/felipedavid/vrcursos/src/infrastructure/repository"
)

type CreateStudentInput struct {
	Name string `json:"name"`
}

type UpdateStudentInput struct {
	Name string `json:"name"`
}

type StudentUsecase interface {
	CreateStudent(ctx context.Context, input CreateStudentInput) (*model.Student, error)
	GetStudent(ctx context.Context, id int) (*model.Student, error)
	GetStudents(ctx context.Context, search string) ([]*model.Student, error)
	UpdateStudent(ctx context.Context, id int, input UpdateStudentInput) (*model.Student, error)
	DeleteStudent(ctx context.Context, id int) error
}

type studentUsecase struct {
	studentRepository repository.IStudentRepository
}

func NewStudentUsecase(repo repository.IStudentRepository) StudentUsecase {
	return &studentUsecase{
		studentRepository: repo,
	}
}

func (u *studentUsecase) CreateStudent(ctx context.Context, input CreateStudentInput) (*model.Student, error) {
	student := &model.Student{
		Name: input.Name,
	}

	err := u.studentRepository.Save(ctx, student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (u *studentUsecase) GetStudent(ctx context.Context, id int) (*model.Student, error) {
	student, err := u.studentRepository.GetStudent(ctx, id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (u *studentUsecase) UpdateStudent(ctx context.Context, id int, input UpdateStudentInput) (*model.Student, error) {
	student, err := u.studentRepository.GetStudent(ctx, id)
	if err != nil {
		return nil, err
	}

	student.Name = input.Name

	err = u.studentRepository.UpdateStudent(ctx, student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (u *studentUsecase) DeleteStudent(ctx context.Context, id int) error {
	err := u.studentRepository.DeleteStudent(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *studentUsecase) GetStudents(ctx context.Context, search string) ([]*model.Student, error) {
	students, err := u.studentRepository.GetStudents(ctx, search)
	if err != nil {
		return nil, err
	}

	if len(students) == 0 {
		if search != "" {
			return nil, domain.ErrStudentNotFound
		}

		return []*model.Student{}, nil
	}

	return students, nil
}
