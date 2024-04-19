package biz

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func PongHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
