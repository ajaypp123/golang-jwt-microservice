package services

import (
	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"
)

type ClientStreamingService struct {
	pb.UnimplementedClientStreamingServiceServer
}

func (s *ClientStreamingService) Average(stream pb.ClientStreamingService_AverageServer) error {
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		sum += req.GetNumber()
		count++
	}
	average := float32(sum) / float32(count)
	return stream.SendAndClose(&pb.AverageResponse{Average: average})
}
