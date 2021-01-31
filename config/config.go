package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Config contains environment variables for the service.
type Config struct {
	Port               string `envconfig:"PORT" default:"8000"`
	Endpoint           string `envconfig:"ENDPOINT" required:"true"`
	MaximumConcurrency int    `envconfig:"MAXIMUMCONCURRENCY" default:"3"`
}

// LoadConfig loads environment variables.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found")
	}

	var c Config

	err := envconfig.Process("", &c)

	return &c, err
}
