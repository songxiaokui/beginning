# *-* coding: utf-8 -*-
from langchain_openai import ChatOpenAI


def new_deepseek_llm(model_name: str = "deepseek-r1") -> ChatOpenAI:
    """
    创建一个深度搜索的LLM模型
    """
    mode_name = "deepseek-chat"
    openai_api_base = "https://api.deepseek.com/v1"
    openai_api_key = ""
    return ChatOpenAI(model_name=model_name, temperature=0.0, openai_api_base=openai_api_base,
                      openai_api_key=openai_api_key)
