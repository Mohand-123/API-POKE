package controllers

import (
	"api_poke/models"
	"api_poke/services"
	"fmt"
	"log"
)

func GetTracks() []models.Track {
	track, err := services.SearchTrack("La bandite")
	if err != nil {
		log.Printf("Erreur lors de la recherche de 'La bandite': %v", err)
		return getFallbackTracks()
	}

	albumCover := ""
	if len(track.Album.Images) > 0 {
		albumCover = track.Album.Images[0].URL
	}

	artistName := "JUL"
	if len(track.Artists) > 0 {
		artistName = track.Artists[0].Name
	}

	duration := fmt.Sprintf("%d:%02d",
		track.Duration/60000,
		(track.Duration%60000)/1000,
	)

	var tracks []models.Track
	tracks = append(tracks, models.Track{
		ID:          track.ID.String(),
		Name:        track.Name,
		Artist:      artistName,
		Duration:    duration,
		Album:       track.Album.Name,
		AlbumCover:  albumCover,
		URI:         string(track.URI),
		ExternalURL: track.ExternalURLs["spotify"],
		PreviewURL:  track.PreviewURL,
	})

	return tracks
}

func getFallbackTracks() []models.Track {
	return []models.Track{
		{Name: "La bandite", Artist: "JUL", Duration: "3:45", Album: "Album 2"},
		{Name: "Autre track", Artist: "JUL", Duration: "4:20", Album: "Album 1"},
	}
}
