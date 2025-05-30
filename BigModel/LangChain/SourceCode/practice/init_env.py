from dotenv import load_dotenv
import os
api_key = "DEEPSEEK_API_KEY"
silicon_flow_key = "SILICON_API_KEY"


def init_env():
    # print("当前工作目录是：", os.getcwd())
    state = load_dotenv(dotenv_path=".env")
    if not state:
        print("API key settings failed. Please check the .env file.")
        return
    print("API_KEY: ", os.getenv("DEEPSEEK_API_KEY"))
    print("API_KEY: ", os.getenv("SILICON_API_KEY"))


if __name__ == '__main__':
    init_env()
