package routers

import (
	"api_poke/controllers"
	"html/template"
	"net/http"
	"path/filepath"
)

func RegisterAlbumsRoutes() {
	http.HandleFunc("/albums", func(w http.ResponseWriter, r *http.Request) {
		data := struct{ Albums interface{} }{Albums: controllers.GetAlbums()}
		templatePath := filepath.Join("..", "..", "templates")
		t := template.Must(template.ParseGlob(filepath.Join(templatePath, "*.html")))
		err := t.ExecuteTemplate(w, "api_poke_albums", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
