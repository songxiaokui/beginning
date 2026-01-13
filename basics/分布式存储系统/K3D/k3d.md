### k3d介绍

> k3d 是一个将轻量级 Kubernetes（k3s）运行在 Docker 容器中的工具，主要用于 本地开发、测试和学习 Kubernetes 环境

### k3d 特点

| 特性       | 说明                                           |
|----------|----------------------------------------------|
| 轻量快速     | 启动一个 Kubernetes 集群只需几秒钟                      |
| 资源占用低    | 基于 k3s，整体比 kubeadm 更轻量，适合本地（如 Mac M1/M2）使用   |
| 多节点支持    | 支持多个 master 和 worker 节点，每个节点运行在一个 Docker 容器中 |
| 容器化部署    | 所有节点均运行在 Docker 中，安装和清理操作简单                  |
| 方便集成     | 支持常用工具如 kubectl、Helm、Ingress、存储组件等           |
| 支持 ARM64 | 可在 Apple Silicon 芯片（M1/M2）上稳定运行              |

### 安装

[下载地址](https://github.com/k3d-io/k3d/releases)
> HOMEBREW_NO_AUTO_UPDATE=1 brew install k3d  
> k3d version
> 

### 创建集群
```shell
k3d cluster create test-cluster  --port "8002:80@loadbalancer" --agents 2 
kubectl get nodes

# 删除集群
k3d cluster delete test-cluster
```