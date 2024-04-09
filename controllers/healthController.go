package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, string([]byte("Healthy")))
	}
}
