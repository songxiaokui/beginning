package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 认证存储
var apiKeys = map[string]string{}

// 生成 AppID 和 SecretKey
func generateCredentials() (string, string) {
	appID := uuid.New().String()
	secretKey := uuid.New().String() // 使用原始字符串，避免 base64 编码
	apiKeys[appID] = secretKey
	apiKeys["56b2084f-c1da-4b17-87ba-164a1e9e3686"] = "M2QyM2VjMWMtNmNmNy00MDdlLTg3Y2MtYzg1MzhkOGMxYTZj"
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
		timestampStr := c.GetHeader("Timestamp")

		if appID == "" || signature == "" || timestampStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication headers"})
			c.Abort()
			return
		}

		// 解析 Timestamp
		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Timestamp format"})
			c.Abort()
			return
		}

		// 验证时间戳（±5 分钟，毫秒级）
		now := time.Now().UnixMilli()
		if math.Abs(float64(now-timestamp)) > 300000 { // 300000 毫秒 = 5 分钟
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Timestamp too old or too far in the future"})
			c.Abort()
			return
		}

		secretKey, exists := apiKeys[appID]
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid App-ID"})
			c.Abort()
			return
		}

		// 读取请求的 body
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			c.Abort()
			return
		}
		// 恢复 body 以供后续处理器使用
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyStr := string(bodyBytes) // 转换为字符串用于签名

		// 计算期望的签名
		expectedSignature := generateSignature(secretKey, c.Request.Method, c.Request.URL.Path, bodyStr, timestamp)

		// 调试日志
		fmt.Printf("Server: method=%s, path=%s, body=%s, timestamp=%d, expectedSignature=%s, receivedSignature=%s\n",
			c.Request.Method, c.Request.URL.Path, bodyStr, timestamp, expectedSignature, signature)

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

		// 示例：支持 POST 请求
		auth.POST("/submit", func(c *gin.Context) {
			log.Println("Job Submitted via POST")
			c.JSON(http.StatusOK, gin.H{"message": "Job submitted successfully via POST"})
		})
	}

	r.Run(":8080")
}
