package routes

import (
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	RegisterPokemonRoutes(mux)
	RegisterErrorRoutes(mux)

	fileServer := http.FileServer(http.Dir("../assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fileServer))

	return mux
}
