# client 单例模式
from langchain_openai import ChatOpenAI
from init_env import init_env, api_key
import os

__client_instance: ChatOpenAI | None = None

init_env()


def get_client() -> ChatOpenAI:
    global __client_instance
    if __client_instance is None:
        __client_instance = ChatOpenAI(
            openai_api_key=os.getenv(api_key),
            openai_api_base="https://api.deepseek.com/v1",
            model_name="deepseek-chat",
        )
    return __client_instance
