package logic

import (
	"context"
	"deepseek_access/internal/consts"
	"deepseek_access/internal/entity"
)

// LLMClient 通用LLM客户端接口
type LLMClient interface {
	Generate(ctx context.Context, prompt string) (*entity.APIResponse, error)
	StreamGenerate(ctx context.Context, prompt string) <-chan entity.StreamResponse
}

func NewLLMClient(config *entity.APIConfig) LLMClient {
	switch config.Type {
	case consts.Ollama:
		return entity.NewOllamaLocalClient(*config)
	case consts.DeepSeek:
		return entity.NewDeepSeekOnlineClient(*config)
	case consts.OpenAI:
		panic("not impl")

	}
	return nil
}
