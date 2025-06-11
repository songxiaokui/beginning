### 使用Django重构后端
1. 安装Django
```shell
pip install django PyMySQL
```

2. 创建项目
```shell
django-admin startproject health_manager
```

3. 创建app
```shell
cd health_manager
python manage.py startapp checkin
```

4. 创建模型升级表
```shell
# models.py 创建模型
# settings 注册应用
python manage.py makemigrations
python manage.py migrate
```

5. 编写业务逻辑
```shell
# 注意模块化时 模块的名称不要写错 要用 models views 否则无法注册
# 编写视图层代码
# 创建路由层 urls.py
# 注册当前模块的路由 health_manager/urls.py
```

6. 部署
```shell
# 使用gunicorn 部署多进程
```

7. 后台管理
```shell
python manage.py createsuperuser
# admin
# sxk123456
# 访问 http://127.0.0.1:5011/admin
```

8. 格式化代码
```shell
pip install black
```

9. 安装langchain相关的大模型库
```shell
pip install langchain langchain-community langchain-openai  openai langgraph langchain_deepseek chroma  
```