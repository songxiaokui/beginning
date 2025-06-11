# *-* coding: utf-8 -*-
from langchain_openai import ChatOpenAI
import os


def new_deepseek_llm():
    """
    创建一个深度搜索的LLM模型
    """
    mode_name = "deepseek-chat"
    openai_api_base = "https://api.deepseek.com/v1"
    openai_api_key = os.getenv("DEEPSEEK_API_KEY", None)

    if openai_api_key is None:
        return None

    return ChatOpenAI(
        model_name=mode_name,
        temperature=0.0,
        openai_api_base=openai_api_base,
        openai_api_key=openai_api_key)
