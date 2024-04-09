package routers

import (
	controllers "github.com/ajaypp123/golang-jwt-microservice/controllers"
	middleware "github.com/ajaypp123/golang-jwt-microservice/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.NewAuthMiddleware().Authenticate)
	incomingRoutes.GET("/users/:user_id", controllers.NewUserController().GetUser())
}
