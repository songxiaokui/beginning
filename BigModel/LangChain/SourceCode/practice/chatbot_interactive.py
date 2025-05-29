# 交互式聊天
# 使用LangGraph持久化上下文
# LangSmith 检查链和代理

"""
1. 安装依赖
pip install langchain-core langgraph>0.2.27

"""
import getpass
import os
from langchain_deepseek import ChatDeepSeek
from practice.init_env import init_env, api_key
from langchain_openai import ChatOpenAI
from langchain.schema import HumanMessage, AIMessage
import os
from langgraph.checkpoint.memory import MemorySaver
from langgraph.graph import START, MessagesState, StateGraph
import asyncio

init_env()


def client() -> ChatOpenAI:
    _chat = ChatOpenAI(
        openai_api_key=os.getenv("DEEPSEEK_API_KEY"),
        openai_api_base="https://api.deepseek.com/v1",
        model_name="deepseek-chat",
    )

    return _chat


def manual_memory() -> None:
    # 调用返回
    response = client().invoke([
        HumanMessage(content="你好啊，我是晓奎？"),
        AIMessage(
            content="你好啊，晓奎！😊 很高兴认识你～今天有什么想聊的，或者需要帮忙的吗？无论是闲聊、问题解答，还是随便聊聊日常，我都在这里哦！✨"),
        HumanMessage(content="你知道我叫啥吗？")
    ])

    print(response.content)


# 消息持久化
# LangGraph 内置了持久层 支持多轮对话 实现状态保存
# 数据持久化支持SQLite和PostgresSQL
def message_persistence(_client: ChatOpenAI, messages: list[HumanMessage]) -> None:
    return None


async def call_model(state: MessagesState):
    return {"messages": await client().ainvoke(state["messages"])}


def call_model2(state: MessagesState):
    return {"messages": client().invoke(state["messages"])}


def new_app():
    workflow = StateGraph(state_schema=MessagesState)
    workflow.add_edge(START, "model")
    workflow.add_node("model", call_model)
    storage = MemorySaver()
    return workflow.compile(checkpointer=storage)


def new_human_message(content: str) -> dict:
    return {"messages": [HumanMessage(content)]}


def new_config(uid: str) -> dict:
    return {"configurable": {"thread_id": uid}}


async def chat():
    app = new_app()
    config = new_config("001")

    # 初始化状态
    output = await app.ainvoke(new_human_message("你好啊，我是晓奎？"), config=config)
    output["messages"][-1].pretty_print()

    output = await app.ainvoke(new_human_message("你现在知道我叫啥了吗？"), config=config)
    output["messages"][-1].pretty_print()

    # 创建一个新的对话 直接问名字
    config2 = new_config("002")
    output = await app.ainvoke(new_human_message("你现在知道我叫啥了吗？"), config=config2)
    output["messages"][-1].pretty_print()
    return


if __name__ == '__main__':
    asyncio.run(chat())
