package app

import (
	"fmt"
	"github.com/22Fariz22/urlcutter/internal/config"
	"github.com/22Fariz22/urlcutter/internal/interceptors"
	"github.com/22Fariz22/urlcutter/internal/url"
	service "github.com/22Fariz22/urlcutter/internal/url/delivery/grpc"
	"github.com/22Fariz22/urlcutter/internal/url/repository"
	"github.com/22Fariz22/urlcutter/pkg/logger"
	"github.com/22Fariz22/urlcutter/pkg/postgres"
	urlcutter "github.com/22Fariz22/urlcutter/proto"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

func (a *App) Run() error {
	log := logger.New("debug")
	im := interceptors.NewInterceptorManager(log, a.cfg)

	l, err := net.Listen("tcp", ":5001")
	if err != nil {
		return err
	}
	defer l.Close()

	server := grpc.NewServer(
		grpc.UnaryInterceptor(im.Logger),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		),
	)

	gRPCServer := service.NewServerGRPC(log, a.cfg, a.UC)
	urlcutter.RegisterURLCutterServiceServer(server, gRPCServer)

	go func() {
		log.Info("Server is listening on port: %v", ":5001")
		if err := server.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-quit
	server.GracefulStop()
	log.Info("Server Exited Properly")

	return nil
}
