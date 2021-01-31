package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/optimusprime77/go-concurrency/client/pigeon"
	"github.com/optimusprime77/go-concurrency/config"
	log "github.com/sirupsen/logrus"
)

// Server contains all the necessary HTTP server attributes.
type Server struct {
	Config *config.Config
	Pigeon *pigeon.Client
	HTTP   *http.Server
	Router *mux.Router
}

// Create creates a new HTTP Server.
func (s *Server) Create(ctx context.Context, config *config.Config) error {

	var pigeonClient pigeon.Client

	if err := pigeonClient.Init(config); err != nil {
		return err
	}

	s.Pigeon = &pigeonClient
	s.Config = config
	s.Router = mux.NewRouter()
	s.HTTP = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Config.Port),
		Handler: s.Router,
	}

	s.setupRoutes()

	return nil
}

// Serve listens on TCP address and serves incoming requests.
func (s *Server) Serve(ctx context.Context) error {

	go func(ctx context.Context, s *http.Server) {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop

		log.Info("Shutdown signal received")

		if err := s.Shutdown(ctx); err != nil {
			log.Error(err.Error())
		}
	}(ctx, s.HTTP)

	log.Infof("Ready at: %s", s.Config.Port)

	if err := s.HTTP.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf(err.Error())
	}

	return nil
}
