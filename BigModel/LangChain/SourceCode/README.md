### LangChain 学习源代码

### 安装依赖
```shell
# 进入虚拟环境
source .venv/bin/activate
# 安装langchain
pip install langchain langchain-community langchain-openai

# 安装openai
pip install openai
# 安装向量数据库
pip install chroma
```

### 安装Notebook
```shell
pip install notebook ipykernel

# 终端启动
jupyter notebook

# 访问
http://localhost:8888/tree
```


### 申请OpenAI Key
```shell
# https://platform.openai.com/api-keys

# oai.furryapi.org

# siliconflow.cn
```

### 使用DeepSeek模型
```shell
# 安装deepseek的相关依赖包
pip install langchain_deepseek   
# 配置deepseek的key https://platform.deepseek.com/api_keys

```

### 环境搭建
```shell
make init
```