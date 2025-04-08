from mcp.server.fastmcp import FastMCP
import httpx
from dotenv import load_dotenv
import os

load_dotenv()


def get_env_config():
    config = {
        'API_KEY': os.getenv('API_KEY'),
        'HOST': os.getenv('HOST', '0.0.0.0'),
        'PORT': int(os.getenv('PORT', 8882))
    }
    return config


mcp = FastMCP("weather", description="实时天气查询服务", host=get_env_config().get("HOST", "0.0.0.0"),
              port=get_env_config().get("PORT", 8882))


@mcp.tool()
def query_weather(city_code: str) -> str:
    """查询指定城市的实时天气。
    Args:
        city_code: 城市名称编码，例如 'Beijing' 110000
    Returns:
        天气信息的字符串描述
    """
    api_key = get_env_config().get("API_KEY")
    url = f"https://restapi.amap.com/v3/weather/weatherInfo?key={api_key}&city={city_code}&extensions=all"
    try:
        response = httpx.get(url)
        response.raise_for_status()
        data = response.json()

        # 高德地图 API 返回格式调整
        forecast = data["forecasts"][0]
        city_name = forecast["city"]
        weather_info = forecast["casts"][0]
        temp_day = weather_info["daytemp"]
        description = weather_info["dayweather"]
        return f"{city_name} 的天气：{description}，白天温度 {temp_day}°C"
    except httpx.HTTPError as e:
        return f"查询失败：{str(e)}"
    except (KeyError, IndexError):
        return "查询失败：无法解析天气数据"


if __name__ == "__main__":
    mcp.run(transport="sse")
