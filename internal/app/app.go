package app

import (
	"fmt"
	"github.com/22Fariz22/urlcutter/internal/config"
	"github.com/22Fariz22/urlcutter/internal/url"
	"github.com/22Fariz22/urlcutter/internal/url/repository"
	"github.com/22Fariz22/urlcutter/pkg/postgres"
	"log"
	"net/http"
)

type App struct {
	cfg        *config.Config
	httpServer *http.Server
	UC         url.UseCase
}

func NewApp(cfg *config.Config) *App {
	if cfg.DatabaseURI == "" {
		//in-memory
		inMemory := repository.NewMemory()

		return &App{
			cfg:        cfg,
			httpServer: nil,
			UC:         inMemory,
		}
	} else {
		// Repository
		db, err := postgres.New(cfg.DatabaseURI, postgres.MaxPoolSize(2))
		if err != nil {
			log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
		}

		repo := repository.NewPGRepository(db)
		return &App{
			cfg:        cfg,
			httpServer: nil,
			UC:         repo,
		}
	}
}

func (a *App) Run() {
	//l := logger.New("debug")

}
