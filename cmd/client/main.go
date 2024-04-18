package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"

	"google.golang.org/grpc"
)

type grpcConn struct {
	conn *grpc.ClientConn
}

func (g *grpcConn) chatServiceCall() {
	// Create a gRPC client
	client := pb.NewChatServiceClient(g.conn)

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

func (g *grpcConn) unaryServiceCall() {
	// Create a gRPC client
	client := pb.NewUnaryServiceClient(g.conn)
	req := &pb.HelloRequest{Name: "Alice"}
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}
	log.Printf("Response: %s", res.GetMessage())
}

func (g *grpcConn) serverStrimingServiceCall() {
	client := pb.NewStreamingServiceClient(g.conn)

	req := &pb.CountRequest{StartNumber: 1}
	stream, err := client.CountNumbers(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call CountNumbers: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("Received number: %d", res.GetNumber())
	}
}

func (g *grpcConn) clientStreamingServiceCall() {
	client := pb.NewClientStreamingServiceClient(g.conn)

	stream, err := client.Average(context.Background())
	if err != nil {
		log.Fatalf("failed to call Average: %v", err)
	}

	numbers := []int32{3, 5, 7, 11, 13}
	for _, num := range numbers {
		req := &pb.NumberRequest{Number: num}
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send number: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive response: %v", err)
	}

	log.Printf("Average: %f", res.GetAverage())
}

func (g *grpcConn) bidirectionalStreamingServiceCall() {
	client := pb.NewBidirectionalServiceClient(g.conn)

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("failed to open stream: %v", err)
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("failed to receive message: %v", err)
			}
			log.Printf("Received message from server: %s", res.GetText())
		}
	}()

	messages := []string{"Hello", "How are you?", "Goodbye"}
	for _, msg := range messages {
		req := &pb.Message{Text: msg}
		if err := stream.Send(req); err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
		time.Sleep(time.Second)
	}

	// Close the client stream
	stream.CloseSend()

	// Wait for the server to finish processing
	time.Sleep(2 * time.Second)
}

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

	grpc_conn := grpcConn{conn: conn}
	var num int

	fmt.Println("Enter a number: \n1. chat service \n2. Unary RPC \n3. Server Streaming RPC " +
		"\n4. Client Streaming RPC \n5. Bidirectional Streaming RPC")

	// Read the number input from the console
	_, err = fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch num {
	case 1:
		grpc_conn.chatServiceCall()
	case 2:
		grpc_conn.unaryServiceCall()
	case 3:
		grpc_conn.serverStrimingServiceCall()
	case 4:
		grpc_conn.clientStreamingServiceCall()
	case 5:
		grpc_conn.bidirectionalStreamingServiceCall()
	default:
		fmt.Println("Invalid option...")
	}
}
