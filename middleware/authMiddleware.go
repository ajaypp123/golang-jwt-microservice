package middleware

import (
	"net/http"

	"github.com/ajaypp123/golang-jwt-microservice/helpers"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Authenticate(c *gin.Context)
}

type AuthMiddlewareImpl struct{}

func NewAuthMiddleware() AuthMiddleware {
	return &AuthMiddlewareImpl{}
}

func (a *AuthMiddlewareImpl) Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No Auth"})
		c.Abort()
		return
	}
	claims, err := helpers.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		c.Abort()
		return
	}
	c.Set("email", claims.Email)
	c.Set("user_id", claims.Uid)
	c.Next()
}
