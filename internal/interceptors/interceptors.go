package interceptors

import (
	"context"

	"github.com/22Fariz22/urlcutter/internal/config"
	"github.com/22Fariz22/urlcutter/pkg/logger"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// InterceptorManager struct
type InterceptorManager struct {
	logger logger.Interface
	cfg    *config.Config
}

// NewInterceptorManager InterceptorManager constructor
func NewInterceptorManager(logger logger.Interface, cfg *config.Config) *InterceptorManager {
	return &InterceptorManager{logger: logger, cfg: cfg}
}

// Logger Interceptor
func (im *InterceptorManager) Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	im.logger.Info("Method: %s, Time: %v, Metadata: %v, Err: %v", info.FullMethod, time.Since(start), md, err)

	return reply, err
}
