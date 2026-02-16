package routes

import (
	"net/http"

	"guide/controllers"
)

func RegisterPokemonRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.HomeHandler)
	mux.HandleFunc("/pokemons", controllers.ListPokemonsHandler)
	mux.HandleFunc("/pokemons/", controllers.PokemonDetailsHandler)
	mux.HandleFunc("/favorites", controllers.FavoritesPageHandler)
	mux.HandleFunc("/favorites/add", controllers.AddFavoriteHandler)
	mux.HandleFunc("/favorites/remove", controllers.RemoveFavoriteHandler)
	mux.HandleFunc("/categories", controllers.CategoriesHandler)
	mux.HandleFunc("/search", controllers.SearchPageHandler)
	mux.HandleFunc("/about", controllers.AboutHandler)
}
