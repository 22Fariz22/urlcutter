package grpc

import (
	"github.com/22Fariz22/urlcutter/internal/config"
	"github.com/22Fariz22/urlcutter/internal/url"
	"github.com/22Fariz22/urlcutter/pkg/logger"
	pb "github.com/22Fariz22/urlcutter/proto"
)

type service struct {
	pb.UnimplementedURLCutterServiceServer
	l   logger.Interface
	cfg *config.Config
	UC  url.UseCase
}

// NewServerGRPC grpc service constructor
func NewServerGRPC(logger logger.Interface, cfg *config.Config, uc url.UseCase) *service {
	return &service{
		UnimplementedURLCutterServiceServer: pb.UnimplementedURLCutterServiceServer{},
		l:                                   logger,
		cfg:                                 cfg,
		UC:                                  uc,
	}
}
