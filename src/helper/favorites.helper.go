package helper

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type FavoriteEntry struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	AddedAt time.Time `json:"addedAt"`
}

type FavoritesData struct {
	Favorites   []FavoriteEntry `json:"favorites"`
	LastUpdated time.Time       `json:"lastUpdated"`
}

const favoritesFile = "favoris.json"

// getFavoritesFilePath retourne le chemin absolu du fichier favoris.json
func getFavoritesFilePath() string {
	dir, _ := os.Getwd()
	// Remonter au dossier parent si on est dans src/
	if filepath.Base(dir) == "src" {
		dir = filepath.Dir(dir)
	}
	return filepath.Join(dir, favoritesFile)
}

// loadFavoritesData charge les données depuis le fichier JSON
func loadFavoritesData() (*FavoritesData, error) {
	filePath := getFavoritesFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		// Si le fichier n'existe pas, retourner une structure vide
		if os.IsNotExist(err) {
			return &FavoritesData{
				Favorites:   []FavoriteEntry{},
				LastUpdated: time.Now(),
			}, nil
		}
		return nil, err
	}

	var favData FavoritesData
	if err := json.Unmarshal(data, &favData); err != nil {
		return nil, err
	}

	return &favData, nil
}

// saveFavoritesData sauvegarde les données dans le fichier JSON
func saveFavoritesData(favData *FavoritesData) error {
	filePath := getFavoritesFilePath()

	favData.LastUpdated = time.Now()

	data, err := json.MarshalIndent(favData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// GetFavorites récupère la liste des IDs favoris depuis le fichier JSON
func GetFavorites(r *http.Request) []int {
	favData, err := loadFavoritesData()
	if err != nil {
		return []int{}
	}

	var favorites []int
	for _, fav := range favData.Favorites {
		favorites = append(favorites, fav.ID)
	}

	return favorites
}

// IsFavorite vérifie si un Pokémon est dans les favoris
func IsFavorite(r *http.Request, pokemonId int) bool {
	favorites := GetFavorites(r)
	for _, id := range favorites {
		if id == pokemonId {
			return true
		}
	}
	return false
}

// AddFavorite ajoute un Pokémon aux favoris
func AddFavorite(w http.ResponseWriter, r *http.Request, pokemonId int) {
	favData, err := loadFavoritesData()
	if err != nil {
		return
	}

	// Vérifier si déjà en favori
	for _, fav := range favData.Favorites {
		if fav.ID == pokemonId {
			return
		}
	}

	// Ajouter le nouveau favori
	newFav := FavoriteEntry{
		ID:      pokemonId,
		Name:    "",
		AddedAt: time.Now(),
	}
	favData.Favorites = append(favData.Favorites, newFav)

	saveFavoritesData(favData)
}

// RemoveFavorite retire un Pokémon des favoris
func RemoveFavorite(w http.ResponseWriter, r *http.Request, pokemonId int) {
	favData, err := loadFavoritesData()
	if err != nil {
		return
	}

	var newFavorites []FavoriteEntry
	for _, fav := range favData.Favorites {
		if fav.ID != pokemonId {
			newFavorites = append(newFavorites, fav)
		}
	}

	favData.Favorites = newFavorites
	saveFavoritesData(favData)
}
