package api

import (
	"encoding/json"
	"go-http/server/api/spotify"
	"log"
	"net/http"

	"github.com/markbates/goth/gothic"
)

// TODO: change auth token requests to be parsed via Authorization header instead of req.Body
// TODO: change non-auth body data to be transmitted via uri instead of http body

func (s *Server) spotifyRoutes() {
	s.HandleFunc("/auth/{provider}/callback", s.getAuthCallback()).Methods("GET")
	s.HandleFunc("/auth/{provider}", s.beginAuth()).Methods("GET")
	s.HandleFunc("/logout/{provider}", s.logout()).Methods("GET")
	s.HandleFunc("/{provider}/playlists", s.playlists()).Methods("GET")
	s.HandleFunc("/{provider}/current-track", s.currentTrack()).Methods("GET")
	s.HandleFunc("/{provider}/add-current-to-playlist", s.addCurrentToPlaylist()).Methods("POST")
	s.HandleFunc("/{provider}/start-track", s.startTrack()).Methods("PUT")
	s.HandleFunc("/{provider}/pause-track", s.pauseTrack()).Methods("PUT")
}

func (s *Server) getAuthCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("Init auth callback")
		newLogIn := spotify.GetAccessTok(w, r)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(newLogIn); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// http.Redirect(w, r, "http://127.0.0.1:3000", http.StatusFound)
	}
}

func (s *Server) beginAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("Init auth")
		spotify.ReqAuth(w, r)
	}
}

func (s *Server) logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func (s *Server) startTrack() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		newRequest := spotify.SpotifyReq{}
		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		spotify.Play(newRequest.AccessToken)
	}
}

func (s *Server) pauseTrack() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		newRequest := spotify.SpotifyReq{}
		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		spotify.Pause(newRequest.AccessToken)
	}
}

func (s *Server) currentTrack() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		newRequest := spotify.SpotifyReq{}
		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		track := spotify.GetCurrentTrack(newRequest.AccessToken)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(track); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) playlists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		newRequest := spotify.SpotifyReq{}
		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		playlists := spotify.GetPlaylists(newRequest.AccessToken)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(playlists); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *Server) addCurrentToPlaylist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		newRequest := spotify.SpotifyReq{}
		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		spotify.UpdatePlaylistWithCurrent(newRequest.AccessToken, newRequest.ID)
	}
}
