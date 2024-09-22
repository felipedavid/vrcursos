package routes

import (
	"net/http"

	"github.com/felipedavid/vrcursos/src/application/controllers"
)

func DefineRoutes(userControllers *controllers.StudentController, courseControllers *controllers.CourseController) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /students", userControllers.StudentList)
	mux.HandleFunc("GET /students/{id}", userControllers.StudentGet)
	mux.HandleFunc("POST /students", userControllers.StudentCreate)
	mux.HandleFunc("PUT /students/{id}", userControllers.StudentUpdate)
	mux.HandleFunc("DELETE /students/{id}", userControllers.StudentDelete)

	mux.HandleFunc("GET /courses", courseControllers.CourseList)
	mux.HandleFunc("GET /courses/{id}", courseControllers.CourseGet)
	mux.HandleFunc("POST /courses", courseControllers.CourseCreate)
	mux.HandleFunc("PUT /courses/{id}", courseControllers.CourseUpdate)
	mux.HandleFunc("DELETE /courses/{id}", courseControllers.CourseDelete)

	return mux
}
