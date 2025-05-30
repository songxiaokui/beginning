# 1. 构建阶段
FROM node:20 AS builder

# 设置工作目录
WORKDIR /app

# 拷贝项目文件
COPY . .

# 安装依赖
RUN npm install

# 构建生产环境代码
RUN npm run build

# 2. 生产镜像阶段
FROM nginx:alpine

# 删除默认 nginx 配置
RUN rm -rf /usr/share/nginx/html/*

# 拷贝打包好的前端静态资源到 nginx 目录
COPY --from=builder /app/dist /usr/share/nginx/html

# 可选：增加自定义 nginx.conf（支持 history 路由等），否则默认配置即可
# COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

# 启动 nginx
CMD ["nginx", "-g", "daemon off;"]
