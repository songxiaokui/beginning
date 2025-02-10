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

// 发送请求
func (c *Client) sendRequest(method, path string, body []byte) (string, error) {
	url := c.BaseURL + path
	timestamp := time.Now().Unix()
	signature := generateSignature(c.SecretKey, method, path, string(body), timestamp)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("App-ID", c.AppID)
	req.Header.Set("Signature", signature)
	req.Header.Set("Timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	return string(respBody), nil
}

// SubmitJob 提交任务
func (c *Client) SubmitJob() (string, error) {
	return c.sendRequest("GET", "/api/submit", nil)
}

// CancelJob 取消任务
func (c *Client) CancelJob() (string, error) {
	return c.sendRequest("GET", "/api/cancel", nil)
}
