package controllers

import (
	"api_poke/models"
	"api_poke/services"
	"log"

	"github.com/zmb3/spotify/v2"
)

func GetAlbums() []models.Album {
	artistID := services.SPOTIFY_ARTIST_ID

	spotifyAlbums, err := services.GetArtistAlbums(spotify.ID(artistID))
	if err != nil {
		log.Printf("Erreur lors de la récupération des albums: %v", err)
		return getFallbackAlbums()
	}

	var albums []models.Album
	for _, album := range spotifyAlbums {
		coverURL := ""
		if len(album.Images) > 0 {
			coverURL = album.Images[0].URL
		}

		releaseYear := 0
		if album.ReleaseDateTime().Year() > 0 {
			releaseYear = album.ReleaseDateTime().Year()
		}

		artistName := "JUL"
		if len(album.Artists) > 0 {
			artistName = album.Artists[0].Name
		}

		albums = append(albums, models.Album{
			ID:          album.ID.String(),
			Title:       album.Name,
			Artist:      artistName,
			Year:        releaseYear,
			Cover:       coverURL,
			URI:         string(album.URI),
			ExternalURL: album.ExternalURLs["spotify"],
		})
	}

	return albums
}

func getFallbackAlbums() []models.Album {
	return []models.Album{
		{Title: "La Machine", Artist: "JUL", Year: 2020, Cover: "/static/covers/la-machine.jpg"},
		{Title: "L'Ovni", Artist: "JUL", Year: 2016, Cover: "/static/covers/lovni.jpg"},
		{Title: "Cœur Blanc", Artist: "JUL", Year: 2022, Cover: "/static/covers/coeur-blanc.jpg"},
	}
}
