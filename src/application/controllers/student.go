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

type StudentController struct {
	studentUsecase usecase.StudentUsecase
}

func NewStudentController(repo repository.IStudentRepository) *StudentController {
	return &StudentController{
		studentUsecase: usecase.NewStudentUsecase(repo),
	}
}

func (c *StudentController) StudentGet(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid id in url")
		return
	}

	student, err := c.studentUsecase.GetStudent(context.Background(), id)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, student, nil)
}

func (c *StudentController) StudentCreate(res http.ResponseWriter, req *http.Request) {
	input := usecase.CreateStudentInput{}
	err := helper.ReadJSON(res, req, &input)
	if err != nil {
		fmt.Fprintf(res, "error reading request body")
		return
	}

	_, err = c.studentUsecase.CreateStudent(context.Background(), input)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusCreated, map[string]any{"message": "student created"}, nil)
}

func (c *StudentController) StudentList(res http.ResponseWriter, req *http.Request) {
	students, err := c.studentUsecase.GetStudents(context.Background())
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, students, nil)
}

func (c *StudentController) StudentUpdate(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid id in url")
		return
	}

	input := usecase.UpdateStudentInput{}
	err = helper.ReadJSON(res, req, &input)
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid request body")
		return
	}

	student, err := c.studentUsecase.UpdateStudent(context.Background(), id, input)
	if err != nil {
		helper.WriteJSON(res, http.StatusInternalServerError, nil, nil)
		return
	}

	helper.WriteJSON(res, http.StatusOK, student, nil)
}

func (c *StudentController) StudentDelete(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		helper.MessageResponse(res, req, http.StatusBadRequest, "invalid id in url")
		return
	}

	err = c.studentUsecase.DeleteStudent(context.Background(), id)
	if err != nil {
		helper.MessageResponse(res, req, http.StatusInternalServerError, "unable to delete student")
		return
	}

	helper.MessageResponse(res, req, http.StatusOK, "student deleted")
}
