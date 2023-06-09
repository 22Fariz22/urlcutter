package grpc

import (
	"context"
	"fmt"

	"github.com/22Fariz22/urlcutter/pkg/grpcerrors"
	pb "github.com/22Fariz22/urlcutter/proto"
	gonanoid "github.com/matoous/go-nanoid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Post(ctx context.Context, url *pb.LongURL) (*pb.ShortURL, error) {
	fmt.Println("here Post handler")

	id, err := gonanoid.Nanoid(10)
	if err != nil {
		s.l.Error("error in short url generator:", err)
		return nil, status.Error(codes.NotFound, "id was not generate")
	}

	short, err := s.UC.Save(ctx, url.LongURL, s.cfg.BaseURL+"/"+id)
	if err != nil {
		if err == grpcerrors.ErrURLExists {
			fmt.Println("handler: already exists:", short)
			return &pb.ShortURL{ShortURL: short}, nil
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	shortURL := short

	return &pb.ShortURL{ShortURL: shortURL}, nil
}

func (s *service) Get(ctx context.Context, url *pb.ShortURL) (*pb.LongURL, error) {
	fmt.Println("here Get  handler", url.ShortURL)

	long, err := s.UC.Get(ctx, url.ShortURL)
	if err != nil {
		s.l.Error("error in handler Get():", err)
		if err == grpcerrors.ErrDoesNotExist {
			return nil, status.Errorf(codes.NotFound, "url does not exists")
		}
		return nil, status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "UC.Get: %v", err)
	}

	return &pb.LongURL{LongURL: long}, nil
}
