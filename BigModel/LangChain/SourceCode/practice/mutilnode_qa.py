"""
多路径执行智能体

"""
from typing import TypedDict
from langchain.schema import BaseMessage, AIMessage, HumanMessage
from langgraph.graph import START, StateGraph
from langgraph.checkpoint.memory import MemorySaver
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
import asyncio
from  singleton_client import get_client


class MutilAgentState(TypedDict):
    messages: list[BaseMessage]
    qa_response: str  # QA
    emotion_response: str  # 情感


async def qa_model(state: MutilAgentState):
    # 使用提示词工程
    prompt = await new_knowledge_prompt_message().ainvoke(state)

    response = await get_client().ainvoke(prompt)
    return {"qa_response": response.content}


async def emotion_model(state: MutilAgentState):
    # 使用提示词工程
    prompt = await new_emotion_prompt_message().ainvoke(state)

    response = await get_client().ainvoke(prompt)
    return {"emotion_response": response.content}


def merge_results(state: MutilAgentState):
    reply = f"""🧠 百度百科回复：{state['qa_response']} \n😊 情感分析：{state['emotion_response']}"""
    return {
        "messages": state["messages"] + [AIMessage(content=reply)]
    }


def new_emotion_prompt_message() -> ChatPromptTemplate:
    return ChatPromptTemplate.from_messages(
        [
            ("system",
             "你现在是一个经验丰富的情感大师，只可以灵活的处理用户问题中只关于情感的部分话题，除了情感话题，你啥也处理不了，也不要去处理非情感话题"),
            MessagesPlaceholder(variable_name="messages"),
        ]
    )


def new_knowledge_prompt_message() -> ChatPromptTemplate:
    return ChatPromptTemplate.from_messages(
        [
            ("system",
             "你现在是百科全书，不要去处理人类情感问题，这里水太深，小老弟你把握不住，你就老老实实处理历史、哲学、以及非情感的所有问题，记住，你不擅长情感问题，你不要回复关于情感的问题"),
            MessagesPlaceholder(variable_name="messages"),
        ]
    )


def new_mutil_node_app():
    # 通过node实现多个agent调用
    workflow = StateGraph(state_schema=MutilAgentState)
    workflow.add_node("qa_model", qa_model)  # 更换为其他的智能体
    workflow.add_node("emo_model", emotion_model)  # 更换为其他的智能体
    workflow.add_node("merge_model", merge_results)

    workflow.add_edge(START, "qa_model")
    workflow.add_edge(START, "emo_model")
    workflow.add_edge("qa_model", "merge_model")
    workflow.add_edge("emo_model", "merge_model")

    return workflow.compile(checkpointer=MemorySaver())


async def run():
    state = {
        "messages": [HumanMessage("我今天很郁闷，同时想知道秦始皇是谁")]
    }
    app = new_mutil_node_app()
    output = await app.ainvoke(state, config={"configurable": {"thread_id": "multi_001"}})
    print(output["messages"][-1].content)


if __name__ == '__main__':
    asyncio.run(run())
