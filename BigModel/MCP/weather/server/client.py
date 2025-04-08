import asyncio
import httpx
from fastapi import FastAPI, HTTPException
from mcp import ClientSession
from mcp.client.sse import sse_client
import uvicorn
from dotenv import load_dotenv
import os

# 加载环境变量
load_dotenv()
API_KEY = os.getenv("API_KEY")

app = FastAPI(title="天气查询MCP客户端")
MCP_SERVER_URL = "http://127.0.0.1:8882/sse"  # SSE 端点


async def get_city_code(city_name: str) -> str:
    """通过高德行政区划查询 API 将城市名称转换为城市代码"""
    url = f"https://restapi.amap.com/v3/config/district?key={API_KEY}&keywords={city_name}&subdistrict=0"
    async with httpx.AsyncClient() as client:
        response = await client.get(url)
        response.raise_for_status()
        data = response.json()
        if data["status"] == "1" and data["districts"]:
            return data["districts"][0]["adcode"]
        else:
            raise ValueError(f"无法找到城市 {city_name} 的代码")


async def query_weather_from_mcp(city_name: str) -> str:
    """查询天气，先将城市名称转换为代码，再调用 MCP 服务端"""
    try:
        # 获取城市代码
        city_code = await get_city_code(city_name)
        print(f"城市 {city_name} 的代码: {city_code}")

        # 调用 MCP 服务端
        async with sse_client(MCP_SERVER_URL) as (reader, writer):
            async with ClientSession(reader, writer) as session:
                await session.initialize()
                result = await session.call_tool("query_weather", {"city_code": city_code})
                print(f"result 类型: {type(result)}")
                print(f"result 内容: {result}")
                if result.content and len(result.content) > 0:
                    return result.content[0].text
                else:
                    raise ValueError("未收到有效天气数据")
    except Exception as e:
        raise ValueError(f"查询天气失败: {str(e)}")


@app.get("/weather/{city}")
async def get_weather(city: str):
    try:
        weather_result = await query_weather_from_mcp(city)
        return {"city": city, "weather": weather_result}
    except Exception as e:
        print(f"错误详情：{str(e)}")
        raise HTTPException(status_code=500, detail=f"查询失败：{str(e)}")


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8080)
