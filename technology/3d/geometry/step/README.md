### 使用conda构建隔离沙盒环境
1. 官方地址
    > https://docs.anaconda.com/free/miniconda/miniconda-other-installer-links/ </br>
2. 基本使用
- 创建一个隔离环境  
  > conda create -n stepenv
- 查询所有虚拟环境  
  > conda env list
- 激活指定的虚拟环境  
  > conda activate stepenv
- 退出当前虚拟环境  
  > conda deactivate
- 查询虚拟环境路径  
  > conda info --envs
- 安装包  
  > conda install python numpy matplotlib
- 注意  
  > conda安装的包不要和pip混淆，尽量使用conda安装
3. 安装pythonocc-core核心包
   > conda install pythonocc-core
4. 选择性激活虚拟环境即可