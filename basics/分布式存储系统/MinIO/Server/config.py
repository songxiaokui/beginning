from pydantic_settings import BaseSettings
from dotenv import load_dotenv

load_dotenv()


class Settings(BaseSettings):
    Endpoint: str
    AccessKey: str
    SecretKey: str
    Secure: bool
    Debug: bool

    class Config:
        env_file = "config/.env"
        env_file_encoding = "utf-8"


settings = Settings()
