# *-* coding: utf-8 -*-
from langchain_openai import ChatOpenAI
import os


def new_qwen_llm():
    """
    创建一个千问的LLM模型
    """
    mode_name = "Qwen/Qwen3-14B"
    openai_api_base = "https://api.siliconflow.cn/v1"
    openai_api_key = os.getenv("QWEN_API_KEY")

    if openai_api_key is None:
        return None

    return ChatOpenAI(
        model_name=mode_name,
        temperature=0.0,
        openai_api_base=openai_api_base,
        openai_api_key=openai_api_key)
