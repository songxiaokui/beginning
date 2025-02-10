package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 认证存储
var apiKeys = map[string]string{}

// 生成 AppID 和 SecretKey
func generateCredentials() (string, string) {
	appID := uuid.New().String()
	secretKey := base64.StdEncoding.EncodeToString([]byte(uuid.New().String()))
	apiKeys[appID] = secretKey // 存储 AppID 和 SecretKey
	return appID, secretKey
}

// 计算 HMAC 签名
func generateSignature(secretKey, method, path, body string, timestamp int64) string {
	data := fmt.Sprintf("%s|%s|%s|%d", method, path, body, timestamp)
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// 认证中间件
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appID := c.GetHeader("App-ID")
		signature := c.GetHeader("Signature")
		timestamp := c.GetHeader("Timestamp")

		if appID == "" || signature == "" || timestamp == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication headers"})
			c.Abort()
			return
		}

		secretKey, exists := apiKeys[appID]
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid App-ID"})
			c.Abort()
			return
		}

		// 计算期望的签名
		expectedSignature := generateSignature(secretKey, c.Request.Method, c.Request.URL.Path, "", time.Now().Unix())

		if signature != expectedSignature {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Signature"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.RedirectTrailingSlash = true

	// 生成 AppID 和 SecretKey
	appID, secretKey := generateCredentials()
	log.Printf("AppID: %s, SecretKey: %s\n", appID, secretKey)

	// 受保护 API
	auth := r.Group("/api")
	auth.Use(authMiddleware())
	{
		auth.GET("/submit", func(c *gin.Context) {
			log.Println("Job Submitted")
			c.JSON(http.StatusOK, gin.H{"message": "Job submitted successfully"})
		})

		auth.GET("/cancel", func(c *gin.Context) {
			log.Println("Job Canceled")
			c.JSON(http.StatusOK, gin.H{"message": "Job canceled successfully"})
		})
	}

	r.Run(":8080")
}
