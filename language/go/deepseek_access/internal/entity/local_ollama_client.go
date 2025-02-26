package entity

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type OllamaLocalClient struct {
	config APIConfig
	client *http.Client
}

func NewOllamaLocalClient(cfg APIConfig) *OllamaLocalClient {
	return &OllamaLocalClient{
		config: cfg,
		client: &http.Client{Timeout: 300 * time.Second},
	}
}

func (c *OllamaLocalClient) Generate(ctx context.Context, prompt string) (*APIResponse, error) {
	start := time.Now()

	reqBody := map[string]interface{}{
		"model":  c.config.ModelName,
		"prompt": prompt,
		"stream": false,
	}

	reqBytes, _ := json.Marshal(reqBody)

	resp, err := c.client.Post(
		c.config.APIBaseURL+"/api/generate",
		"application/json",
		bytes.NewReader(reqBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var result struct {
		Response string `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &APIResponse{
		Content: result.Response,
		Latency: time.Since(start),
	}, nil
}

// StreamGenerate Ollama本地客户端流式实现
func (c *OllamaLocalClient) StreamGenerate(ctx context.Context, prompt string) <-chan StreamResponse {
	ch := make(chan StreamResponse)

	go func() {
		defer close(ch)

		reqBody := map[string]interface{}{
			"model":  c.config.ModelName,
			"prompt": prompt,
			"stream": true,
		}

		reqBytes, _ := json.Marshal(reqBody)

		resp, err := c.client.Post(
			c.config.APIBaseURL+"/api/generate",
			"application/json",
			bytes.NewReader(reqBytes),
		)
		if err != nil {
			ch <- StreamResponse{Error: fmt.Errorf("请求失败: %w", err)}
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			ch <- StreamResponse{Error: fmt.Errorf("API返回错误状态码: %d", resp.StatusCode)}
			return
		}

		decoder := json.NewDecoder(resp.Body)
		for {
			var result struct {
				Response string `json:"response"`
				Done     bool   `json:"done"`
				Error    string `json:"error"`
			}

			if err := decoder.Decode(&result); err != nil {
				if err == io.EOF {
					return
				}
				ch <- StreamResponse{Error: fmt.Errorf("解析失败: %w", err)}
				return
			}

			if result.Error != "" {
				ch <- StreamResponse{Error: fmt.Errorf("API错误: %s", result.Error)}
				return
			}

			if result.Response != "" {
				ch <- StreamResponse{Content: result.Response}
			}

			if result.Done {
				return
			}
		}
	}()

	return ch
}
