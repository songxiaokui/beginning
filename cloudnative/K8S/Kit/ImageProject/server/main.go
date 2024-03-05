package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", Pong)
	r.GET("/health", Health)
	r.Run("0.0.0.0:8881") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
