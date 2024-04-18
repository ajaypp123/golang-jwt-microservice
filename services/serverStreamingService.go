package services

import (
	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"
)

type ServerStreamingService struct {
	pb.UnimplementedStreamingServiceServer
}

func (s *ServerStreamingService) CountNumbers(req *pb.CountRequest, stream pb.StreamingService_CountNumbersServer) error {
	startNumber := req.GetStartNumber()
	for i := startNumber; i < startNumber+10; i++ {
		response := &pb.CountResponse{Number: i}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
	return nil
}
