from langchain_deepseek import ChatDeepSeek
from practice.init_env import init_env, api_key

from langchain.schema import HumanMessage
from singleton_client import get_client
import os


def three_package_call():
    init_env()
    llm = ChatDeepSeek(
        temperature=0,
        max_tokens=None,
        timeout=None,
        max_retries=2,
        api_key=os.getenv(api_key),
        model="deepseek-chat",
    )
    response = llm.invoke("你是谁？")
    print(response.content)
    # response
    # 我是DeepSeek Chat，由深度求索公司创造的智能AI助手！🤖✨ 我的使命是帮助你解答问题、提供信息、陪你聊天，甚至帮你处理各种文本和文件。无论是学习、工作，还是日常生活中的疑问，都可以来问我！有什么我可以帮你的吗？😊


def standard_call():
    response = get_client().invoke([
        HumanMessage(content="你好啊，我是武安君？")
    ])

    print(response.content)

    # 流式返回
    data_stream = get_client().stream([
        HumanMessage(content="给我介绍一下战国时期的武安君，字数控制在100字之内")
    ])

    buffer = ""
    for chunk in data_stream:
        content = chunk.content or ""
        buffer += content

        # 输出策略：遇到标点或者每 N 个字刷新一次
        if any(p in content for p in "，。：；！？\n") or len(buffer) > 6:
            print(buffer, end="", flush=True)
            buffer = ""

    # 输出剩余部分
    if buffer:
        print(buffer, end="", flush=True)


if __name__ == '__main__':
    standard_call()
