package entity

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// DeepSeekOnlineClient 线上DeepSeek客户端实现
type DeepSeekOnlineClient struct {
	config APIConfig
	client *http.Client
}

func NewDeepSeekOnlineClient(cfg APIConfig) *DeepSeekOnlineClient {
	return &DeepSeekOnlineClient{
		config: cfg,
		client: &http.Client{},
	}
}

func (c *DeepSeekOnlineClient) Generate(ctx context.Context, prompt string) (*APIResponse, error) {
	start := time.Now()

	reqBody := map[string]interface{}{
		"model": c.config.ModelName,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}
	fmt.Println(reqBody)

	reqBytes, _ := json.Marshal(reqBody)
	fullURL := fmt.Sprintf("%s/chat/completions", strings.TrimSuffix(c.config.APIBaseURL, "/"))
	req, _ := http.NewRequest("POST", fullURL, bytes.NewReader(reqBytes))
	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result.Choices) == 0 {
		return nil, fmt.Errorf("no content in response")
	}

	return &APIResponse{
		Content: result.Choices[0].Message.Content,
		Latency: time.Since(start),
	}, nil
}

func (c *DeepSeekOnlineClient) StreamGenerate(ctx context.Context, prompt string) <-chan StreamResponse {
	ch := make(chan StreamResponse)

	go func() {
		defer close(ch)

		fullURL := fmt.Sprintf("%s/chat/completions", strings.TrimSuffix(c.config.APIBaseURL, "/"))

		reqBody := map[string]interface{}{
			"model": c.config.ModelName,
			"messages": []map[string]string{
				{"role": "user", "content": prompt},
			},
			"stream": true,
		}

		reqBytes, _ := json.Marshal(reqBody)

		req, _ := http.NewRequest("POST", fullURL, bytes.NewReader(reqBytes))
		req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "text/event-stream")

		resp, err := c.client.Do(req)
		if err != nil {
			ch <- StreamResponse{Error: err}
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			ch <- StreamResponse{Error: fmt.Errorf("API错误 %d: %s", resp.StatusCode, body)}
			return
		}

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				ch <- StreamResponse{Error: err}
				return
			}

			rawData := bytes.TrimSpace(line)
			if len(rawData) == 0 {
				continue
			}

			// 处理SSE协议格式
			if bytes.HasPrefix(rawData, []byte("data: ")) {
				jsonData := bytes.TrimPrefix(rawData, []byte("data: "))

				// 添加结束标记处理
				if bytes.Equal(jsonData, []byte("[DONE]")) {
					return // 正常结束
				}

				var event struct {
					Choices []struct {
						Delta struct {
							Content string `json:"content"`
						} `json:"delta"`
					} `json:"choices"`
					Error struct {
						Message string `json:"message"`
					} `json:"error"`
				}

				if err := json.Unmarshal(jsonData, &event); err != nil {
					// 添加更友好的错误提示
					log.Printf("解析警告: 跳过无效数据块 | 原始数据: %s", string(jsonData))
					continue
				}

				if event.Error.Message != "" {
					ch <- StreamResponse{Error: fmt.Errorf(event.Error.Message)}
					return
				}

				if len(event.Choices) > 0 && event.Choices[0].Delta.Content != "" {
					ch <- StreamResponse{Content: event.Choices[0].Delta.Content}
				}
			}
		}
	}()
	return ch
}
