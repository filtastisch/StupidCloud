package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	router := chi.NewRouter()

	return &Server{
		server: &http.Server{
			Addr:    ":443",
			Handler: router,
		},
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.server.Close()
}
