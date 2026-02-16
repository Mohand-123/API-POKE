package services

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	client *spotify.Client
	auth   *spotifyauth.Authenticator
)

const (
	SPOTIFY_CLIENT_ID     = "b4252c29687a4faf9835a3ec0a4e247c"
	SPOTIFY_CLIENT_SECRET = "db5c872cf2c74181ab71bdb9bb27d59c"
	SPOTIFY_ARTIST_ID     = "3IW7ScrzXmPvZhB27hmfgy"
)

func InitSpotifyClient() error {
	clientID := SPOTIFY_CLIENT_ID
	clientSecret := SPOTIFY_CLIENT_SECRET

	if clientID == "your_client_id_here" || clientSecret == "your_client_secret_here" {
		return fmt.Errorf("veuillez configurer vos credentials Spotify dans services/spotify.go")
	}

	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return fmt.Errorf("erreur lors de l'obtention du token: %v", err)
	}

	httpClient := spotifyauth.New().Client(context.Background(), token)
	client = spotify.New(httpClient)

	log.Println("✅ Client Spotify initialisé avec succès")

	// Tester la recherche de JUL pour obtenir le bon ID
	artist, err := SearchArtistByName("JUL")
	if err != nil {
		log.Printf("⚠️  Erreur recherche JUL: %v", err)
	} else {
		log.Printf("✅ JUL trouvé - ID: %s, Nom: %s", artist.ID, artist.Name)
	}

	return nil
}

func SearchArtistByName(artistName string) (*spotify.FullArtist, error) {
	if client == nil {
		return nil, fmt.Errorf("client Spotify non initialisé")
	}

	ctx := context.Background()
	results, err := client.Search(ctx, artistName, spotify.SearchTypeArtist)
	if err != nil {
		return nil, err
	}

	if results.Artists == nil || len(results.Artists.Artists) == 0 {
		return nil, fmt.Errorf("aucun artiste trouvé pour: %s", artistName)
	}

	return &results.Artists.Artists[0], nil
}

func GetClient() *spotify.Client {
	return client
}

func SearchArtist(artistName string) (*spotify.FullArtist, error) {
	if client == nil {
		return nil, fmt.Errorf("client Spotify non initialisé")
	}

	ctx := context.Background()
	results, err := client.Search(ctx, artistName, spotify.SearchTypeArtist)
	if err != nil {
		return nil, err
	}

	if results.Artists == nil || len(results.Artists.Artists) == 0 {
		return nil, fmt.Errorf("aucun artiste trouvé pour: %s", artistName)
	}

	return &results.Artists.Artists[0], nil
}

func GetArtistAlbums(artistID spotify.ID) ([]spotify.SimpleAlbum, error) {
	if client == nil {
		return nil, fmt.Errorf("client Spotify non initialisé")
	}

	ctx := context.Background()
	albums, err := client.GetArtistAlbums(ctx, artistID, []spotify.AlbumType{spotify.AlbumTypeAlbum})
	if err != nil {
		return nil, err
	}

	return albums.Albums, nil
}

func GetArtistTopTracks(artistID spotify.ID, country string) ([]spotify.FullTrack, error) {
	if client == nil {
		return nil, fmt.Errorf("client Spotify non initialisé")
	}

	ctx := context.Background()
	tracks, err := client.GetArtistsTopTracks(ctx, artistID, country)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func GetTrack(trackID spotify.ID) (*spotify.FullTrack, error) {
	if client == nil {
		return nil, fmt.Errorf("client Spotify non initialisé")
	}

	ctx := context.Background()
	track, err := client.GetTrack(ctx, trackID)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func GetAlbum(albumID spotify.ID) (*spotify.FullAlbum, error) {
	if client == nil {
		return nil, fmt.Errorf("client Spotify non initialisé")
	}

	ctx := context.Background()
	album, err := client.GetAlbum(ctx, albumID)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func SearchTrack(trackName string) (*spotify.FullTrack, error) {
	if client == nil {
		return nil, fmt.Errorf("client Spotify non initialisé")
	}

	ctx := context.Background()
	results, err := client.Search(ctx, trackName+" JUL", spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	}

	if results.Tracks == nil || len(results.Tracks.Tracks) == 0 {
		return nil, fmt.Errorf("aucun morceau trouvé pour: %s", trackName)
	}

	return &results.Tracks.Tracks[0], nil
}
