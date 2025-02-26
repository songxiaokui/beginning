package svc

import (
	"deepseek_access/internal/entity"
	"deepseek_access/internal/logic"
)

type ServiceContext struct {
	Config    entity.APIConfig
	LLMClient logic.LLMClient
}

func NewServiceContext(cf entity.APIConfig) *ServiceContext {
	return &ServiceContext{
		Config:    cf,
		LLMClient: logic.NewLLMClient(&cf),
	}
}
