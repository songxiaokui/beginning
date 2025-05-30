# 使用官方 Python 镜像作为基础镜像
FROM python:3.11-slim

# 设置工作目录
WORKDIR /app

# 拷贝代码文件
COPY . /app
RUN mkdir -pv db

# 安装依赖（如果未来加 requirements.txt 可替换）
RUN pip install -r requirements.txt

# 初始化数据库（可选：首次启动时用）
# CMD 中调用 Python 脚本时已自动执行 init_db()

# 暴露端口
EXPOSE 5001

# 启动 Flask 应用（关闭 debug）
CMD ["python", "cmd/main.py"]
