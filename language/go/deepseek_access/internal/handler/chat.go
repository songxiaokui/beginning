package handler

import (
	"context"
	"deepseek_access/internal/svc"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ChatHandler struct {
	ServiceContext *svc.ServiceContext
}

func NewChatHandler(svcCtx *svc.ServiceContext) *ChatHandler {
	return &ChatHandler{
		ServiceContext: svcCtx,
	}
}

type ChatRequest struct {
	Message string `json:"message" binding:"required"`
	Stream  bool   `json:"stream"`
}

// Chat 普通聊天接口
func (h *ChatHandler) Chat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.ServiceContext.LLMClient.Generate(c.Request.Context(), req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": resp.Content,
		"latency": resp.Latency.String(),
	})
}

// StreamChat 流式聊天接口
func (h *ChatHandler) StreamChat(c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 设置流式响应头
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	streamCh := h.ServiceContext.LLMClient.StreamGenerate(c.Request.Context(), req.Message)

	// 创建超时上下文
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Minute)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		case resp, ok := <-streamCh:
			if !ok {
				c.SSEvent("end", "stream ended")
				return
			}
			if resp.Error != nil {
				c.SSEvent("error", resp.Error.Error())
				return
			}
			c.SSEvent("message", resp.Content)
			c.Writer.Flush()
		}
	}
}
