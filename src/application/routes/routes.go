package routes

import (
	"net/http"

	"github.com/felipedavid/vrcursos/src/application/controllers"
)

func DefineRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.HelloWorld)

	return mux
}
