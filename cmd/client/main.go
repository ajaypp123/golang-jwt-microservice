package main

import (
	"context"
	"log"
	"os"

	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"

	"google.golang.org/grpc"
)

func main() {
	// Start gRPC client
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "9090"
	}

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(":"+grpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewChatServiceClient(conn)

	// Call the SendMessage RPC
	stream, err := client.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("Error calling SendMessage RPC: %v", err)
	}

	// Send a message
	err = stream.Send(&pb.MessageRequest{Sender: "client", Message: "Hello from client"})
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	// Receive a response
	response, err := stream.Recv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	// Print the response
	log.Printf("Received response from server: %v", response)
}
