# *-* coding: utf-8 -*-
from langchain_openai import ChatOpenAI


def new_qwen_llm(model_name: str = "qwen") -> ChatOpenAI:
    """
    创建一个千问的LLM模型
    """
    mode_name = "Qwen/Qwen3-14B"
    openai_api_base = "https://api.siliconflow.cn/v1"
    openai_api_key = ""
    return ChatOpenAI(model_name=model_name, temperature=0.0, openai_api_base=openai_api_base,
                      openai_api_key=openai_api_key)
