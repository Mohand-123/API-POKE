package controllers

import (
	"fmt"
	"guide/helper"
	"guide/services"
	"net/http"
	"strconv"
)

type ListPageData struct {
	Pokemons    []services.Pokemon
	Page        int
	TotalPages  int
	HasPrevious bool
	HasNext     bool
	SearchQuery string
	FilterType  string
	MinWeight   string
	MaxWeight   string
	MinExp      string
	MaxExp      string
	TotalCount  int
	Favorites   []int
}

type PokemonDetailsData struct {
	Pokemon    *services.Pokemon
	IsFavorite bool
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Désactiver le cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	helper.RenderTemplate(w, r, "home", nil)
}

func ListPokemonsHandler(w http.ResponseWriter, r *http.Request) {
	// Désactiver le cache
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	// Récupérer les paramètres de pagination, recherche et filtre
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	searchQuery := r.URL.Query().Get("search")
	filterType := r.URL.Query().Get("type")
	minWeight := r.URL.Query().Get("min_weight")
	maxWeight := r.URL.Query().Get("max_weight")
	minExp := r.URL.Query().Get("min_exp")
	maxExp := r.URL.Query().Get("max_exp")
	perPage := 20

	result, statusCode, serviceErr := services.GetPokemonsWithPagination(page, perPage, searchQuery, filterType, minWeight, maxWeight, minExp, maxExp)
	if statusCode != http.StatusOK || serviceErr != nil {
		http.Error(w, fmt.Sprintf("Erreur service - code :%d \n message: %s", statusCode, serviceErr.Error()), statusCode)
		return
	}

	data := ListPageData{
		Pokemons:    result.Pokemons,
		Page:        result.Page,
		TotalPages:  result.TotalPages,
		HasPrevious: result.HasPrevious,
		HasNext:     result.HasNext,
		SearchQuery: searchQuery,
		FilterType:  filterType,
		MinWeight:   minWeight,
		MaxWeight:   maxWeight,
		MinExp:      minExp,
		MaxExp:      maxExp,
		TotalCount:  result.TotalCount,
		Favorites:   helper.GetFavorites(r),
	}

	helper.RenderTemplate(w, r, "list_pokemons", data)
}

func PokemonDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract Pokemon ID from URL path
	idStr := r.URL.Path[len("/pokemons/"):]
	if idStr == "" {
		http.Redirect(w, r, "/pokemons", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/pokemons", http.StatusSeeOther)
		return
	}

	pokemon, statusCode, serviceErr := services.GetPokemonByID(id)
	if statusCode != http.StatusOK || serviceErr != nil {
		http.Error(w, fmt.Sprintf("Erreur service - code :%d \n message: %s", statusCode, serviceErr.Error()), statusCode)
		return
	}

	data := PokemonDetailsData{
		Pokemon:    pokemon,
		IsFavorite: helper.IsFavorite(r, id),
	}

	helper.RenderTemplate(w, r, "pokemons-details", data)
}

func FavoritesPageHandler(w http.ResponseWriter, r *http.Request) {
	favoriteIds := helper.GetFavorites(r)

	if len(favoriteIds) == 0 {
		helper.RenderTemplate(w, r, "favorites", nil)
		return
	}

	// Récupérer tous les pokémons
	allPokemons, statusCode, err := services.GetAllPokemons()
	if statusCode != http.StatusOK || err != nil {
		http.Error(w, fmt.Sprintf("Erreur service - code :%d \n message: %s", statusCode, err.Error()), statusCode)
		return
	}

	// Filtrer pour ne garder que les favoris
	var favoritePokemons []services.Pokemon
	for _, pokemon := range *allPokemons {
		for _, favId := range favoriteIds {
			if pokemon.Id == favId {
				favoritePokemons = append(favoritePokemons, pokemon)
				break
			}
		}
	}

	data := map[string]interface{}{
		"Pokemons":   favoritePokemons,
		"TotalCount": len(favoritePokemons),
	}

	helper.RenderTemplate(w, r, "favorites", data)
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer le type depuis les paramètres
	filterType := r.URL.Query().Get("type")

	if filterType == "" {
		// Afficher la page de sélection de type
		helper.RenderTemplate(w, r, "categories", nil)
		return
	}

	// Récupérer tous les Pokémons du type spécifié
	result, statusCode, err := services.GetPokemonsWithPagination(1, 1000, "", filterType, "", "", "", "")
	if statusCode != http.StatusOK || err != nil {
		http.Error(w, fmt.Sprintf("Erreur service - code :%d \n message: %s", statusCode, err.Error()), statusCode)
		return
	}

	data := map[string]interface{}{
		"Pokemons":   result.Pokemons,
		"TypeName":   filterType,
		"TotalCount": result.TotalCount,
	}

	helper.RenderTemplate(w, r, "categories", data)
}

func SearchPageHandler(w http.ResponseWriter, r *http.Request) {
	searchQuery := r.URL.Query().Get("q")

	if searchQuery == "" {
		// Afficher la page de recherche vide
		helper.RenderTemplate(w, r, "search", nil)
		return
	}

	// Rechercher les Pokémons
	result, statusCode, err := services.GetPokemonsWithPagination(1, 100, searchQuery, "", "", "", "", "")
	if statusCode != http.StatusOK || err != nil {
		http.Error(w, fmt.Sprintf("Erreur service - code :%d \n message: %s", statusCode, err.Error()), statusCode)
		return
	}

	data := map[string]interface{}{
		"Pokemons":    result.Pokemons,
		"SearchQuery": searchQuery,
		"TotalCount":  result.TotalCount,
	}

	helper.RenderTemplate(w, r, "search", data)
}

func AddFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("pokemon_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	helper.AddFavorite(w, r, id)

	// Rediriger vers la page précédente ou vers les détails
	referer := r.Header.Get("Referer")
	if referer != "" {
		http.Redirect(w, r, referer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, fmt.Sprintf("/pokemons/%d", id), http.StatusSeeOther)
	}
}

func RemoveFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.FormValue("pokemon_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	helper.RemoveFavorite(w, r, id)

	// Rediriger vers la page précédente ou vers la liste
	referer := r.Header.Get("Referer")
	if referer != "" {
		http.Redirect(w, r, referer, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/favorites", http.StatusSeeOther)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, r, "about", nil)
}
