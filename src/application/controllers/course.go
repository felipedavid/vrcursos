package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/felipedavid/vrcursos/src/core/domain/usecase"
	"github.com/felipedavid/vrcursos/src/core/helper"
	"github.com/felipedavid/vrcursos/src/infrastructure/repository"
)

type CourseController struct {
	courseUsecase usecase.CourseUsecase
}

func NewCourseController(repo repository.ICourseRepository) *CourseController {
	return &CourseController{
		courseUsecase: usecase.NewCourseUsecase(repo),
	}
}

func (c *CourseController) CourseGet(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid id in url")
		return
	}

	course, err := c.courseUsecase.GetCourse(context.Background(), id)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, course, nil)
}

func (c *CourseController) CourseCreate(res http.ResponseWriter, req *http.Request) {
	input := usecase.CreateCourseInput{}
	err := helper.ReadJSON(res, req, &input)
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid request body")
		return
	}

	_, err = c.courseUsecase.CreateCourse(context.Background(), input)
	if err != nil {
		helper.MessageResponse(res, req, http.StatusInternalServerError, "internal server error")
		return
	}

	helper.MessageResponse(res, req, http.StatusCreated, "course created")
}

func (c *CourseController) CourseList(res http.ResponseWriter, req *http.Request) {
	courses, err := c.courseUsecase.GetCourses(context.Background())
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, courses, nil)
}

func (c *CourseController) CourseUpdate(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid id in url")
		return
	}

	input := usecase.UpdateCourseInput{}
	err = helper.ReadJSON(res, req, &input)
	if err != nil {
		fmt.Fprintf(res, "error reading request body")
		return
	}

	_, err = c.courseUsecase.UpdateCourse(context.Background(), id, input)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, map[string]any{"message": "course updated"}, nil)
}

func (c *CourseController) CourseDelete(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid id in url")
		return
	}

	err = c.courseUsecase.DeleteCourse(context.Background(), id)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, map[string]any{"message": "course deleted"}, nil)
}

func (c *CourseController) EnrollStudent(res http.ResponseWriter, req *http.Request) {
	studentID, err := strconv.Atoi(req.PathValue("studentID"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid userID in url")
		return
	}

	courseID, err := strconv.Atoi(req.PathValue("courseID"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid courseID in url")
		return
	}

	err = c.courseUsecase.EnrollStudent(context.Background(), courseID, studentID)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.MessageResponse(res, req, http.StatusOK, "student enrolled in the course successfully")
}

func (c *CourseController) UnenrollStudent(res http.ResponseWriter, req *http.Request) {
	studentID, err := strconv.Atoi(req.PathValue("studentID"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid userID in url")
		return
	}

	courseID, err := strconv.Atoi(req.PathValue("courseID"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid courseID in url")
		return
	}

	err = c.courseUsecase.UnenrollStudent(context.Background(), courseID, studentID)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.MessageResponse(res, req, http.StatusOK, "student unenrolled from the course successfully")
}
