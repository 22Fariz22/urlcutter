package main

import (
	"github.com/22Fariz22/urlcutter/internal/app"
	"github.com/22Fariz22/urlcutter/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg)
	app.Run()
}
