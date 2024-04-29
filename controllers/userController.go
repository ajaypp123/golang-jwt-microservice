package controllers

import (
	"net/http"

	"github.com/ajaypp123/golang-jwt-microservice/models"
	services "github.com/ajaypp123/golang-jwt-microservice/services"

	"github.com/gin-gonic/gin"
)

type userController struct{}

func NewUserController() *userController {
	return &userController{}
}

func (u *userController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := services.NewUserService().Signup(c, &user)
		c.JSON(res.Code, res.Message)
	}
}

func (u *userController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := services.NewUserService().Login(c, &user)
		c.JSON(res.Code, res.Message)
	}
}

func (u *userController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		res := services.NewUserService().GetUser(c, &userId)
		c.JSON(res.Code, res.Message)
	}
}
