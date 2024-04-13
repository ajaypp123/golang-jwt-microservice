package services

import (
	"log"

	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"
)

// ChatServiceServer is used to implement ChatServiceServer.
type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
}

// SendMessage implements ChatServiceServer.SendMessage
func (s *ChatServiceServer) SendMessage(stream pb.ChatService_SendMessageServer) error {
	for {
		message, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Received message from %s: %s", message.Sender, message.Message)

		response := &pb.MessageResponse{
			Sender:  message.Sender,
			Message: "Message received",
		}

		if err := stream.Send(response); err != nil {
			return err
		}
	}
}
