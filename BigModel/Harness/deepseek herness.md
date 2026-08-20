## Deepseek Herness使用
1. 准备
   * Node环境：去官网下载 LTS 版本 或者使用 nvm
   * Deepseek API key

2. 安装
   * 命令行安装
   > npx @deepseek-ai/dsh web
   * 源代码安装
   ```shell
   git clone https://github.com/deepseek-ai/deepseek-harness.git
   cd deepseek-harness
   nvm use v25.6.1
   pnpm install 
   pnpm run build
   pnpm dsh web
   ```
3. 配置环境变量(或者内部配置)
   ```shell
   echo "export DEEPSEEK_API_KEY=你的key" >> ~/.bashrc  
   # or
   echo "export DEEPSEEK_API_KEY=你的key" >> ~/.zshrc 
   ```
   
4. 打开 http://127.0.0.1:3080

5. 使用docker-compose启动并配置认证，可实现本地操作服务器  
   * 修改.env中挂载的目录  
   * make restart
   * 访问: http://127.0.0.1:3081