package routes

import (
	"net/http"

	"guide/controllers"
)

func RegisterErrorRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/404", controllers.NotFoundHandler)
}
