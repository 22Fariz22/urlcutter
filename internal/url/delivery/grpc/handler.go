package grpc

import (
	"context"
	"fmt"
	pb "github.com/22Fariz22/urlcutter/proto"
)

func (s *service) Post(ctx context.Context, url *pb.LongURL) (*pb.ShortURL, error) {
	fmt.Println("here Post handler")

	s.UC.Save(ctx)

	return &pb.ShortURL{}, nil
}

func (s *service) Get(ctx context.Context, url *pb.ShortURL) (*pb.LongURL, error) {
	fmt.Println("here Get  handler")

	s.UC.Get(ctx)

	return &pb.LongURL{}, nil
}
