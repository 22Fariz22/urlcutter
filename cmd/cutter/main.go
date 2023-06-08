package main

import (
	"github.com/22Fariz22/urlcutter/internal/app"
	"github.com/22Fariz22/urlcutter/internal/config"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg)
	err := app.Run()
	if err != nil {
		log.Error(err)
		return
	}
}
