package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/biz"
	"server/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.GET("/ping", biz.PongHandle)
	r.GET("/health", biz.HealthHandle)
	r.POST("/upload", biz.UploadHandle) // 用于文件上传的接口
	r.StaticFS("/file", http.Dir("./static"))
	r.Run("0.0.0.0:8881") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
