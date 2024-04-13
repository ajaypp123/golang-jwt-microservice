package main

import (
	"log"
	"net"
	"os"

	"github.com/ajaypp123/golang-jwt-microservice/helpers/logger"
	pb "github.com/ajaypp123/golang-jwt-microservice/pb_generated"
	routers "github.com/ajaypp123/golang-jwt-microservice/routers"
	"github.com/ajaypp123/golang-jwt-microservice/services"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// Start gRPC server
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "9090"
	}

	go func() {
		if err := StartGRPCServer(grpcPort); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// Start HTTP server
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routers.AuthRoutes(router)
	routers.HealthRoutes(router)
	routers.UserRoutes(router)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server, ", err)
	}
	logger.Logger.Info("Started application on 8080 ....")
}

// StartGRPCServer starts the gRPC server
func StartGRPCServer(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &services.ChatServiceServer{})
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}
