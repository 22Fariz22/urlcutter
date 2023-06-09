package grpc

import (
	"context"
	"fmt"
	"github.com/22Fariz22/urlcutter/pkg/grpcerrors"
	pb "github.com/22Fariz22/urlcutter/proto"
	gonanoid "github.com/matoous/go-nanoid"
)

func (s *service) Post(ctx context.Context, url *pb.LongURL) (*pb.ShortURL, error) {
	fmt.Println("here Post handler")

	id, err := gonanoid.Nanoid(10)
	if err != nil {
		s.l.Error("error in short url generator:", err)
		return nil, err
	}

	short, err := s.UC.Save(ctx, url.LongURL, s.cfg.BaseURL+"/"+id)
	if err != nil {
		if err == grpcerrors.ErrURLExists {
			fmt.Println("handler: already exists:", short)
			return &pb.ShortURL{ShortURL: short}, nil //status.Errorf(grpcerrors.ParseGRPCErrStatusCode(err), "UC.Save: %v", err)
		}
		return nil, err
	}

	shortUrl := short

	return &pb.ShortURL{ShortURL: shortUrl}, nil
}

func (s *service) Get(ctx context.Context, url *pb.ShortURL) (*pb.LongURL, error) {
	fmt.Println("here Get  handler", url.ShortURL)

	long, err := s.UC.Get(ctx, url.ShortURL)
	if err != nil {
		s.l.Error("error in handler Get():", err)
		return nil, err
	}

	return &pb.LongURL{LongURL: long}, nil
}
