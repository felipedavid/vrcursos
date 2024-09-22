package usecase

import (
	"context"
	"fmt"

	"github.com/felipedavid/vrcursos/src/core/model"
	"github.com/felipedavid/vrcursos/src/infrastructure/repository"
)

type CreateCourseInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCourseInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetCourseOutput struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	HowManyEnrolled int    `json:"how_many_enrolled"`
}

type CourseUsecase interface {
	CreateCourse(ctx context.Context, input CreateCourseInput) (*model.Course, error)
	GetCourse(ctx context.Context, id int) (*GetCourseOutput, error)
	GetCourses(ctx context.Context) ([]*GetCourseOutput, error)
	UpdateCourse(ctx context.Context, id int, input UpdateCourseInput) (*model.Course, error)
	DeleteCourse(ctx context.Context, id int) error
	EnrollStudent(ctx context.Context, courseID, studentID int) error
	UnenrollStudent(ctx context.Context, courseID, studentID int) error
}

type courseUsecase struct {
	courseRepository  repository.ICourseRepository
	studentRepository repository.IStudentRepository
}

func NewCourseUsecase(courseRepo repository.ICourseRepository, studentRepo repository.IStudentRepository) CourseUsecase {
	return &courseUsecase{
		courseRepository:  courseRepo,
		studentRepository: studentRepo,
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

func (u *courseUsecase) GetCourse(ctx context.Context, id int) (*GetCourseOutput, error) {
	course, err := u.courseRepository.GetCourse(ctx, id)
	if err != nil {
		return nil, err
	}

	nStudentsEnrolled, err := u.courseRepository.HowManyEnrolled(ctx, id)
	if err != nil {
		return nil, err
	}

	courseOutput := &GetCourseOutput{
		ID:              int(course.ID),
		Name:            course.Name,
		Description:     course.Description,
		HowManyEnrolled: nStudentsEnrolled,
	}

	return courseOutput, nil
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

func (u *courseUsecase) GetCourses(ctx context.Context) ([]*GetCourseOutput, error) {
	courses, err := u.courseRepository.GetCourses(ctx)
	if err != nil {
		return nil, err
	}

	var coursesOutput []*GetCourseOutput

	// TODO: Refactor this in only a single query
	for _, course := range courses {
		nStudentsEnrolled, err := u.courseRepository.HowManyEnrolled(ctx, int(course.ID))
		if err != nil {
			return nil, err
		}

		courseOutput := &GetCourseOutput{
			ID:              int(course.ID),
			Name:            course.Name,
			Description:     course.Description,
			HowManyEnrolled: nStudentsEnrolled,
		}

		coursesOutput = append(coursesOutput, courseOutput)
	}

	return coursesOutput, nil
}

type errCourseFull struct {
	MaxStudents int
}

func (e *errCourseFull) Error() string {
	return fmt.Sprintf("course is full, max students: %d", e.MaxStudents)
}

type errEnrolledTooManyCourses struct {
	MaxCourses int
}

func (e *errEnrolledTooManyCourses) Error() string {
	return fmt.Sprintf("student is already enrolled in %d courses", e.MaxCourses)
}

var ErrCourseFull = &errCourseFull{MaxStudents: 10}
var ErrEnrolledTooManyCourses = &errEnrolledTooManyCourses{MaxCourses: 3}
var ErrStudentAlreadyEnrolled = repository.ErrStudentAlreadyEnrolled

func (u *courseUsecase) EnrollStudent(ctx context.Context, courseID, studentID int) error {
	nCourses, err := u.studentRepository.EnrolledInHowManyCourses(ctx, studentID)
	if err != nil {
		return err
	}

	if nCourses >= ErrEnrolledTooManyCourses.MaxCourses {
		return ErrEnrolledTooManyCourses
	}

	nStudentsEnrolled, err := u.courseRepository.HowManyEnrolled(ctx, courseID)
	if err != nil {
		return err
	}

	if nStudentsEnrolled >= ErrCourseFull.MaxStudents {
		return ErrCourseFull
	}

	err = u.courseRepository.AddStudentToCourse(ctx, courseID, studentID)
	if err != nil {
		if err == repository.ErrStudentAlreadyEnrolled {
			return ErrStudentAlreadyEnrolled
		}
		return err
	}

	return nil
}

func (u *courseUsecase) UnenrollStudent(ctx context.Context, courseID, studentID int) error {
	err := u.courseRepository.RemoveStudentFromCourse(ctx, courseID, studentID)
	if err != nil {
		return err
	}

	return nil
}
