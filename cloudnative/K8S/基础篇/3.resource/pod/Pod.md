#### 1. Pod属性
- 凡是调度、网络、存储、安全类似于传统虚拟机级别的设置，都是Pod级别
- 描述的是机器的属性与设置
- 设置机器的存储 - Pod的存储定义
- 设置机器的网卡 - Pod的网络定义
- 设置机器的防火墙 - Pod的安全定义
- 设置服务运行在哪个机器上 - Pod的调度
- 凡事与容器NameSpace相关的设置也是属于Pod级别

#### 2. Pod的重要字段
- NodeSelector 节点选择器
- HostAliases 定义Pod中的hosts文件
- shareProcessNamespace 共享Pid namespace
- ImagePullPolicy 镜像拉取策略
- lifecycle 容器生命周期，当状态发生改变时触发钩子
  - postStart 启动之后触发 如果钩子执行报错 则Pod启动报错
  - preStop 停止之前触发 当收到Kill信号，会先执行，同时会阻塞停止信号，同步执行

#### 3. Projected Volume 投射数据卷
- 
