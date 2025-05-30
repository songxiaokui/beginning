# client 单例模式
from langchain_openai import ChatOpenAI
from init_env import init_env, api_key, silicon_flow_key
import os

__client_instance: ChatOpenAI | None = None
__client_instance2: ChatOpenAI | None = None

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


def get_qwen_client() -> ChatOpenAI:
    global __client_instance2
    if __client_instance2 is None:
        __client_instance2 = ChatOpenAI(
            openai_api_key=os.getenv(silicon_flow_key),
            openai_api_base="https://api.siliconflow.cn/v1",
            model_name="Qwen/Qwen3-14B",
        )
    return __client_instance2
