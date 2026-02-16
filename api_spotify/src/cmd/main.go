package main

import (
	"api_poke/controllers"
	"api_poke/routers"
	"api_poke/services"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func loadTemplates() *template.Template {
	templatePath := filepath.Join("..", "..", "templates")
	tmpl := template.New("")

	// Charger uniquement les templates Spotify
	files := []string{"layout.html", "home.html", "albums.html", "track.html", "error.html"}
	for _, file := range files {
		_, err := tmpl.ParseFiles(filepath.Join(templatePath, file))
		if err != nil {
			log.Printf("Erreur lors du chargement de %s: %v", file, err)
		}
	}
	return tmpl
}

func main() {
	if err := services.InitSpotifyClient(); err != nil {
		log.Printf("⚠️  Erreur lors de l'initialisation du client Spotify: %v", err)
		log.Println("📝 L'application fonctionnera avec des données de fallback")
		log.Println("📝 Pour activer Spotify: modifiez services/spotify.go avec vos credentials")
	}

	tmpl := loadTemplates()

	routers.RegisterAlbumsRoutes()
	routers.RegisterTracksRoutes()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("..", "..", "assets", "css")))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		albums := controllers.GetAlbums()
		tracks := controllers.GetTracks()

		data := struct {
			Albums []struct {
				Title, Artist string
				Year          int
				Cover         string
			}
			Tracks []struct{ Name, Duration string }
		}{
			Albums: make([]struct {
				Title, Artist string
				Year          int
				Cover         string
			}, len(albums)),
			Tracks: make([]struct{ Name, Duration string }, len(tracks)),
		}

		for i, album := range albums {
			data.Albums[i] = struct {
				Title, Artist string
				Year          int
				Cover         string
			}{
				Title:  album.Title,
				Artist: album.Artist,
				Year:   album.Year,
				Cover:  album.Cover,
			}
		}

		limit := len(tracks)
		if limit > 5 {
			limit = 5
		}
		for i := 0; i < limit; i++ {
			data.Tracks[i] = struct{ Name, Duration string }{
				Name:     tracks[i].Name,
				Duration: tracks[i].Duration,
			}
		}

		err := tmpl.ExecuteTemplate(w, "home.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("🚀 Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
