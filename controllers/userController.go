package controllers

import (
	services "github.com/ajaypp123/golang-jwt-microservice/service"

	"github.com/gin-gonic/gin"
)

type userController struct{}

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.NewUserService().Signup(c)
	}
}

func (u *userController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.NewUserService().Login(c)
	}
}

func (u *userController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		services.NewUserService().GetUser(c)
	}
}
