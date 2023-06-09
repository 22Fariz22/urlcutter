package config

import (
	"flag"
	"log"

	"github.com/caarlos0/env/v6"
)

// Config config
type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	BaseURL     string `env:"BASE_URL" envDefault:"http://localhost:8080"`
	DatabaseURI string `env:"DATABASE_URI"`
}

// NewConfig create configuration
func NewConfig() *Config {
	cfg := Config{}

	flag.StringVar(&cfg.RunAddress, "a", "", "server address")

	flag.StringVar(&cfg.DatabaseURI, "d", "", "database address")

	flag.Parse()

	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}

	return &cfg
}
