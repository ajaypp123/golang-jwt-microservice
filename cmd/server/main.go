package main

import (
	"log"
	"os"

	"github.com/ajaypp123/golang-jwt-microservice/helpers/logger"
	routers "github.com/ajaypp123/golang-jwt-microservice/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
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
