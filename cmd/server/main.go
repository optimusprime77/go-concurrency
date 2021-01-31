package main

import (
	"context"

	"github.com/optimusprime77/go-concurrency/config"
	"github.com/optimusprime77/go-concurrency/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	log.Info("Starting ...")

	ctx := context.Background()
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	var s server.Server

	if err := s.Create(ctx, config); err != nil {
		log.Fatal(err.Error())
	}

	if err := s.Serve(ctx); err != nil {
		log.Fatal(err.Error())
	}
}
