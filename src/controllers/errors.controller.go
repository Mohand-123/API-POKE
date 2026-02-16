package controllers

import (
	"guide/helper"
	"guide/models"
	"net/http"
)

// errorDisplay gère l'affichage de la page d'erreur.
// On récupère le code et le message passés en query string /error?code=404&message=Page%30introuvable
func ErrorDisplay(w http.ResponseWriter, r *http.Request) {
	// Récupération "simple" (string) depuis la query. Exemple: /error?code=404&message=Page%30introuvable
	data := models.Error{
		Code:    r.FormValue("code"),    // "404"
		Message: r.FormValue("message"), // "Page introuvable"
	}

	// On délègue le rendu au helper de templates (buffer + redirection si erreur de rendu)
	helper.RenderTemplate(w, r, "error", data)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := models.Error{
		Code:    "404",
		Message: "Page introuvable",
	}
	w.WriteHeader(http.StatusNotFound)
	helper.RenderTemplate(w, r, "error", data)
}
