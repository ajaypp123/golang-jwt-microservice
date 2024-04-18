package services

import (
	"context"

	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"
)

type UnaryService struct {
	pb.UnimplementedUnaryServiceServer
}

func (s *UnaryService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	message := "Hello, " + name
	return &pb.HelloResponse{Message: message}, nil
}
