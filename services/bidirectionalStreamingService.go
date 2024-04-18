package services

import (
	"io"
	"log"

	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"
)

type BidirectionalStreamingService struct {
	pb.UnimplementedBidirectionalServiceServer
}

func (s *BidirectionalStreamingService) Chat(stream pb.BidirectionalService_ChatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", req.GetText())

		res := &pb.Message{Text: "Server received: " + req.GetText()}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}
