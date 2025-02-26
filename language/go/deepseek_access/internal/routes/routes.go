package routes

import (
	"deepseek_access/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *handler.ChatHandler) {
	api := router.Group("/api")
	{
		api.POST("/chat", h.Chat)
		api.POST("/chat/stream", h.StreamChat)
	}
}
