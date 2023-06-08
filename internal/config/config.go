package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	RunAddress  string `env:"RUN_ADDRESS"`
	DatabaseURI string `env:"DATABASE_URI"`
}

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
