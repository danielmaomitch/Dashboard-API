package api

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}

	s.taskRoutes()
	s.spotifyRoutes()

	return s
}
