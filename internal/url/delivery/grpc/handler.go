package grpc

import (
	"context"

	"github.com/22Fariz22/urlcutter/pkg/grpcerrors"
	pb "github.com/22Fariz22/urlcutter/proto"
	gonanoid "github.com/matoous/go-nanoid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Post(ctx context.Context, url *pb.LongURL) (*pb.ShortURL, error) {
	id, err := gonanoid.Nanoid(10)
	if err != nil {
		s.l.Error("error in short url generator:", err)
		return nil, status.Error(codes.NotFound, "id was not generate")
	}

	short, err := s.UC.Save(ctx, s.l, url.LongURL, s.cfg.BaseURL+"/"+id)
	if err != nil {
		if err == grpcerrors.ErrURLExists {
			return &pb.ShortURL{ShortURL: short}, nil
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.ShortURL{ShortURL: short}, nil
}

func (s *service) Get(ctx context.Context, url *pb.ShortURL) (*pb.LongURL, error) {
	long, err := s.UC.Get(ctx, s.l, url.ShortURL)
	if err != nil {
		s.l.Error("error in handler Get():", err)
		if err == grpcerrors.ErrDoesNotExist {
			return nil, status.Errorf(codes.NotFound, "url does not exists")
		}
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "UC.Get: %v", err)
	}

	return &pb.LongURL{LongURL: long}, nil
}
