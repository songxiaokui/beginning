### List MCP tools and description

1. 源代码
```python
import asyncio
from mcp.client.sse import sse_client
from mcp import ClientSession
import argparse


async def main(endpoint: str, tools_name: str = None):
    async with sse_client(endpoint) as (reader, writer):
        async with ClientSession(reader, writer) as session:
            await session.initialize()
            tools_list = await session.list_tools()

            has_exists = False
            for tool in tools_list.tools:
                if tools_name and tools_name != tool.name:
                    continue
                print(f" 工具名称: {tool.name}")
                print(f" 工具描述: {tool.description}")
                print(f" 工具结构化参数: {tool.inputSchema}")
                print("-----" * 5)
                has_exists = True
            if not has_exists:
                print(f"没有找到 {tools_name} 工具")


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description="MCP 工具查看器")
    parser.add_argument("--host", help="MCP Server Host", required=True)
    parser.add_argument("--tool", help="指定要查看的工具名")
    args = parser.parse_args()

    asyncio.run(main(args.host, args.tool))
```

2. 安装依赖
```shell
cat << EOF > requirements.txt
annotated-types==0.7.0
anyio==4.9.0
attrs==25.3.0
certifi==2025.6.15
click==8.2.1
fastapi==0.115.12
h11==0.16.0
httpcore==1.0.9
httpx==0.28.1
httpx-sse==0.4.1
idna==3.10
jsonschema==4.24.0
jsonschema-specifications==2025.4.1
markdown-it-py==3.0.0
mcp==1.10.1
mdurl==0.1.2
pydantic==2.11.7
pydantic-core==2.33.2
pydantic-settings==2.10.1
pygments==2.19.2
python-dotenv==1.1.1
python-multipart==0.0.20
referencing==0.36.2
rich==14.0.0
rpds-py==0.25.1
shellingham==1.5.4
sniffio==1.3.1
sse-starlette==2.3.6
starlette
typer==0.16.0
typing-extensions==4.14.0
typing-inspection==0.4.1
uvicorn==0.35.0
websockets==15.0.1
EOF
```
> uv venv --python 3.12  
> source .venv/bin/activate  
> uv pip install -r requirements.txt  
> uv run python mcp_tool_viewer.py --host http://127.0.0.1:8000