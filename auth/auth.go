package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/spotify"
)

const (
	key = "0q8JFhN2zrvzBFlvx7y46pKfvjjDpb7"
)

func NewAuth() {
	spotifyClientId := os.Getenv("CLIENT_ID")
	spotifySecret := os.Getenv("CLIENT_SECRET")
	scopes := "user-read-email user-read-currently-playing  playlist-read-private " +
		"user-modify-playback-state playlist-modify-private playlist-modify-public"

	store := sessions.NewCookieStore([]byte(key))
	store.Options.HttpOnly = true
	gothic.Store = store
	goth.UseProviders(
		spotify.New(spotifyClientId, spotifySecret, "http://127.0.0.1:8080/auth/spotify/callback", scopes),
	)
}
