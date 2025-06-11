# -*- coding: utf-8 -*-
from langchain_openai import ChatOpenAI


class LLMClientRegister(object):
    __client_maps = dict()

    @classmethod
    def register(cls, model_name: str, client: ChatOpenAI):
        if model_name not in cls.__client_maps:
            cls.__client_maps[model_name] = client

    @classmethod
    def get(cls, model_name: str):
        return cls.__client_maps.get(model_name)


class Singleton(object):
    _instance: LLMClientRegister = None

    @classmethod
    def get_instance(cls):
        if cls._instance is None:
            cls._instance = LLMClientRegister()
        return cls._instance

    def __init__(self):
        if self._instance is not None:
            raise Exception("This class is a singleton!")


singleton = Singleton()
