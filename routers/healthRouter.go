package routers

import (
	"github.com/ajaypp123/golang-jwt-microservice/controllers"
	"github.com/gin-gonic/gin"
)

func HealthRoutes(inncommingRoutes *gin.Engine) {
	inncommingRoutes.GET("/health", controllers.GetHealth())
}
