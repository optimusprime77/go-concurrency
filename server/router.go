package server

import (
	"github.com/optimusprime77/go-concurrency/server/internal/handler"
)

const v1API string = "/api/v1/"

func (s *Server) setupRoutes() {
	s.Router.HandleFunc("/_healthz", handler.Healthz).Methods("GET").Name("Health")
	s.Router.HandleFunc(v1API+"pigeon", handler.Pigeon(s.Pigeon)).Methods("GET").Name("Pigeon")
}
