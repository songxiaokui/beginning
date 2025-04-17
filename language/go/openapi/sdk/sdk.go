package sdk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	AppID     string
	SecretKey string
	BaseURL   string
}

// 计算 HMAC 签名
func generateSignature(secretKey, method, path, body string, timestamp int64) string {
	data := fmt.Sprintf("%s|%s|%s|%d", method, path, body, timestamp)
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// SendRequest 发送请求
func (c *Client) SendRequest(method, path string, body []byte) (string, error) {
	url := c.BaseURL + path
	timestamp := time.Now().UnixMilli() // 使用毫秒级时间戳提高精度
	bodyStr := ""                       // 默认空字符串
	if body != nil {
		bodyStr = string(body) // 只有非 nil 时转换为字符串
	}
	signature := generateSignature(c.SecretKey, method, path, bodyStr, timestamp)

	// 调试日志
	fmt.Printf("Client: method=%s, path=%s, body=%s, timestamp=%d, signature=%s\n", method, path, bodyStr, timestamp, signature)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("create request failed: %v", err)
	}
	req.Header.Set("App-ID", c.AppID)
	req.Header.Set("Signature", signature)
	req.Header.Set("Timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("send request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response body failed: %v", err)
	}
	return string(respBody), nil
}

// SubmitJob 提交任务
func (c *Client) SubmitJob() (string, error) {
	return c.SendRequest("GET", "/api/submit", nil)
}

// CancelJob 取消任务
func (c *Client) CancelJob() (string, error) {
	return c.SendRequest("GET", "/api/cancel", nil)
}
