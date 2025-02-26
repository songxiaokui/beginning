package entity

import (
	"deepseek_access/internal/consts"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

// APIConfig 配置结构体
type APIConfig struct {
	Port       int64            `yaml:"port" json:"port"`
	Type       consts.ModelType `yaml:"type" json:"type"`
	APIBaseURL string           `yaml:"api_base_url" json:"api_base_url"`
	APIKey     string           `yaml:"api_key" json:"api_key"`
	ModelName  string           `yaml:"model_name" json:"model_name"`
	UseLocal   bool             `yaml:"use_local" json:"use_local"`
}

// APIResponse 通用响应结构体
type APIResponse struct {
	Content string
	Latency time.Duration
}

type StreamResponse struct {
	Content string
	Error   error
}

func LoadConfig(configPath string) (APIConfig, error) {
	var config APIConfig
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return config, err
	}
	return config, nil
}
