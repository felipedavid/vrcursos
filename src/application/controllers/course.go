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
	studentUsecase usecase.CourseUsecase
}

func NewCourseController(repo repository.ICourseRepository) *CourseController {
	return &CourseController{
		studentUsecase: usecase.NewCourseUsecase(repo),
	}
}

func (c *CourseController) CourseGet(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid id in url")
		return
	}

	course, err := c.studentUsecase.GetCourse(context.Background(), id)
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
		fmt.Fprintf(res, "error reading request body")
		return
	}

	_, err = c.studentUsecase.CreateCourse(context.Background(), input)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusCreated, map[string]any{"message": "course created"}, nil)
}

func (c *CourseController) CourseList(res http.ResponseWriter, req *http.Request) {

	courses, err := c.studentUsecase.GetCourses(context.Background())
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

	_, err = c.studentUsecase.UpdateCourse(context.Background(), id, input)
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

	err = c.studentUsecase.DeleteCourse(context.Background(), id)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, map[string]any{"message": "course deleted"}, nil)
}
