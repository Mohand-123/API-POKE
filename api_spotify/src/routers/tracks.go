package routers

import (
	"api_poke/controllers"
	"html/template"
	"net/http"
	"path/filepath"
)

func RegisterTracksRoutes() {
	http.HandleFunc("/tracks", func(w http.ResponseWriter, r *http.Request) {
		tracks := controllers.GetTracks()

		data := struct {
			Tracks interface{}
		}{
			Tracks: tracks,
		}

		templatePath := filepath.Join("..", "..", "templates")
		t := template.Must(template.ParseGlob(filepath.Join(templatePath, "*.html")))
		err := t.ExecuteTemplate(w, "api_poke_track", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
