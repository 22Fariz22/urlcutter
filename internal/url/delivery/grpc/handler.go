package grpc

import (
	"context"
	"fmt"
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

	// сделать gRPC error
	short, err := s.UC.Save(ctx, url.LongURL, id)
	if err != nil {
		s.l.Error("error in handler Post():", err)
		//добавить if  errAlreadyExist
		return nil, err
	}

	shortUrl := "http://localhost:8080/" + short

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
