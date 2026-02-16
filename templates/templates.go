package templates

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

// variable globale qui stocke tous les templates parsés
var listTemp *template.Template

// Load charge tous les fichiers de templates au démarrage.
// En prod, pense à utiliser une racine stable (ex: "templates/*.gohtml") ou `embed`.
func Load() {
	listTemplates, errTemplates := template.ParseGlob("./../../templates/*.html")
	if errTemplates != nil {
		// On log l'erreur puis on arrête le programme.
		log.Fatalf("Erreur chargement des templates : %s", errTemplates.Error())
	}
	// On garde le résultat dans la variable globale pour le réutiliser partout.
	listTemp = listTemplates
}

// RenderTemplate exécute un template et écrit la réponse HTTP.
func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	var buffer bytes.Buffer

	// Exécution du template `name` avec les données `data` dans le buffer (pas encore envoyé au client).
	errRender := listTemp.ExecuteTemplate(&buffer, name, data)
	if errRender != nil {
		// Si le rendu échoue, on redirige vers une page d'erreur dédiée.
		// On passe le message via la query string (encodé pour l’URL).
		http.Redirect(
			w,
			r,
			fmt.Sprintf(
				"/error?code=%d&message=%s",
				http.StatusInternalServerError,
				url.QueryEscape("Erreur lors du chargement de la page"),
			),
			http.StatusSeeOther,
		)
		return
	}

	// Si tout est OK, on vide le buffer vers la réponse HTTP.
	_, _ = buffer.WriteTo(w)
}
