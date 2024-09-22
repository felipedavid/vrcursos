package usecase

import (
	"context"

	"github.com/felipedavid/vrcursos/src/core/model"
	"github.com/felipedavid/vrcursos/src/infrastructure/repository"
)

type CreateCourseInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCourseInput struct {
	Name        string `json:"name"`
	Description string `json:"name"`
}

type CourseUsecase interface {
	CreateCourse(ctx context.Context, input CreateCourseInput) (*model.Course, error)
	GetCourse(ctx context.Context, id int) (*model.Course, error)
	GetCourses(ctx context.Context) ([]*model.Course, error)
	UpdateCourse(ctx context.Context, id int, input UpdateCourseInput) (*model.Course, error)
	DeleteCourse(ctx context.Context, id int) error
}

type courseUsecase struct {
	courseRepository repository.ICourseRepository
}

func NewCourseUsecase(repo repository.ICourseRepository) CourseUsecase {
	return &courseUsecase{
		courseRepository: repo,
	}
}

func (u *courseUsecase) CreateCourse(ctx context.Context, input CreateCourseInput) (*model.Course, error) {
	course := &model.Course{
		Name:        input.Name,
		Description: input.Description,
	}

	err := u.courseRepository.Save(ctx, course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (u *courseUsecase) GetCourse(ctx context.Context, id int) (*model.Course, error) {
	course, err := u.courseRepository.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (u *courseUsecase) UpdateCourse(ctx context.Context, id int, input UpdateCourseInput) (*model.Course, error) {
	course, err := u.courseRepository.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	course.Name = input.Name
	course.Description = input.Description

	err = u.courseRepository.UpdateCourse(ctx, course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (u *courseUsecase) DeleteCourse(ctx context.Context, id int) error {
	err := u.courseRepository.DeleteCourse(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *courseUsecase) GetCourses(ctx context.Context) ([]*model.Course, error) {
	courses, err := u.courseRepository.GetCourses(ctx)
	if err != nil {
		return nil, err
	}

	return courses, nil
}
