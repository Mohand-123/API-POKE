package main

import (
	"fmt"
	"guide/helper"
	routers "guide/routes"
	"net/http"
)

func main() {
	helper.Load()
	// Chargement des routes du serveur
	serveRouter := routers.RegisterRoutes()
	// Message d'information indiquant que le serveur est lancé
	fmt.Println("Serveur lancé : http://localhost:8080")
	// Lancement du serveur HTTP sur le port 8080
	http.ListenAndServe("localhost:8080", serveRouter)
}
