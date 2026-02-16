package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Pokemon struct {
	Id             int    `json:"pokedexId"`
	Name           string `json:"name"`
	Image          string `json:"image"`
	Weight         string `json:"weight"`
	BaseExperience int    `json:"base_experience"`
	Types          []struct {
		Name string `json:"name"`
	} `json:"apiTypes"`
}

type PokemonListResult struct {
	Pokemons    []Pokemon
	TotalCount  int
	Page        int
	PerPage     int
	TotalPages  int
	HasPrevious bool
	HasNext     bool
}

func GetAllPokemons() (*[]Pokemon, int, error) {
	_client := http.Client{
		Timeout: time.Second * 60,
	}

	request, requestErr := http.NewRequest(http.MethodGet, "https://pokebuildapi.fr/api/v1/pokemon", nil)
	if requestErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur preparation requete - %s", requestErr.Error())
	}

	response, responseErr := _client.Do(request)
	if responseErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur envois requete - %s", responseErr.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("Erreur response - %s", response.Status)
	}

	var listePokemons []Pokemon

	decodeErr := json.NewDecoder(response.Body).Decode(&listePokemons)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur decode données - %s", decodeErr.Error())
	}

	return &listePokemons, response.StatusCode, nil
}

// GetPokemonsWithPagination retourne une liste paginée de Pokémons
func GetPokemonsWithPagination(page, perPage int, searchQuery, filterType, minWeight, maxWeight, minExp, maxExp string) (*PokemonListResult, int, error) {
	allPokemons, statusCode, err := GetAllPokemons()
	if err != nil {
		return nil, statusCode, err
	}

	// Convertir les filtres de poids et d'expérience
	var minWeightFloat, maxWeightFloat, minExpInt, maxExpInt float64
	var hasMinWeight, hasMaxWeight, hasMinExp, hasMaxExp bool

	if minWeight != "" {
		if val, err := strconv.ParseFloat(minWeight, 64); err == nil {
			minWeightFloat = val
			hasMinWeight = true
		}
	}
	if maxWeight != "" {
		if val, err := strconv.ParseFloat(maxWeight, 64); err == nil {
			maxWeightFloat = val
			hasMaxWeight = true
		}
	}
	if minExp != "" {
		if val, err := strconv.ParseFloat(minExp, 64); err == nil {
			minExpInt = val
			hasMinExp = true
		}
	}
	if maxExp != "" {
		if val, err := strconv.ParseFloat(maxExp, 64); err == nil {
			maxExpInt = val
			hasMaxExp = true
		}
	}

	// Filtrer les Pokémons par recherche et type
	var filteredPokemons []Pokemon
	for _, pokemon := range *allPokemons {
		// Recherche par nom OU par type (FT1 - recherche sur 2 propriétés minimum)
		if searchQuery != "" {
			matchName := strings.Contains(strings.ToLower(pokemon.Name), strings.ToLower(searchQuery))
			matchType := false
			for _, t := range pokemon.Types {
				if strings.Contains(strings.ToLower(t.Name), strings.ToLower(searchQuery)) {
					matchType = true
					break
				}
			}
			// Si la recherche ne correspond ni au nom ni au type, on passe au suivant
			if !matchName && !matchType {
				continue
			}
		}

		// Filtre par type (différent de la recherche)
		if filterType != "" {
			hasType := false
			for _, t := range pokemon.Types {
				if strings.EqualFold(t.Name, filterType) {
					hasType = true
					break
				}
			}
			if !hasType {
				continue
			}
		}

		// Filtre par poids
		if hasMinWeight || hasMaxWeight {
			if weight, err := strconv.ParseFloat(pokemon.Weight, 64); err == nil {
				if hasMinWeight && weight < minWeightFloat {
					continue
				}
				if hasMaxWeight && weight > maxWeightFloat {
					continue
				}
			}
		}

		// Filtre par expérience
		if hasMinExp && float64(pokemon.BaseExperience) < minExpInt {
			continue
		}
		if hasMaxExp && float64(pokemon.BaseExperience) > maxExpInt {
			continue
		}

		filteredPokemons = append(filteredPokemons, pokemon)
	}

	totalCount := len(filteredPokemons)
	totalPages := (totalCount + perPage - 1) / perPage

	if page < 1 {
		page = 1
	}
	if page > totalPages && totalPages > 0 {
		page = totalPages
	}

	// Calculer l'index de début et de fin pour la pagination
	start := (page - 1) * perPage
	end := start + perPage
	if end > totalCount {
		end = totalCount
	}

	var paginatedPokemons []Pokemon
	if start < totalCount {
		paginatedPokemons = filteredPokemons[start:end]
	}

	result := &PokemonListResult{
		Pokemons:    paginatedPokemons,
		TotalCount:  totalCount,
		Page:        page,
		PerPage:     perPage,
		TotalPages:  totalPages,
		HasPrevious: page > 1,
		HasNext:     page < totalPages,
	}

	return result, http.StatusOK, nil
}

// GetPokemonByID retourne un Pokémon spécifique par son ID
func GetPokemonByID(id int) (*Pokemon, int, error) {
	allPokemons, statusCode, err := GetAllPokemons()
	if err != nil {
		return nil, statusCode, err
	}

	for _, pokemon := range *allPokemons {
		if pokemon.Id == id {
			return &pokemon, http.StatusOK, nil
		}
	}

	return nil, http.StatusNotFound, fmt.Errorf("Pokemon non trouvé")
}
