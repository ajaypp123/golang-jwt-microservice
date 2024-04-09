package routers

import (
	"github.com/ajaypp123/golang-jwt-microservice/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(inncommingRoutes *gin.Engine) {
	inncommingRoutes.POST("/users/signup", controllers.NewUserController().Signup())
	inncommingRoutes.POST("/users/login", controllers.NewUserController().Login())
}
