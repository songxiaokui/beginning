# Lightweight CephFS Deployment with Docker Compose

This project provides a simple two-VM setup for deploying CephFS using Docker Compose.

## 🧩 Components

- **VM1 (Server)**: Runs Ceph MON, MGR, MDS, and OSD via Docker Compose.
- **VM2 (Client)**: Mounts CephFS via `mount`.

## 🚀 Quick Start

### 1. On VM1: Setup CephFS Cluster

启动集群:
终端一: 
> make init  

终端二: 
> make up

关闭集群:
> make down

部署单节点集群排错:
[部署排错](./deploy-cephfs-troubleshooting.md)
[客户端挂载排错](./client-mount-troubleshooting.md)

最后在两个不同的机器上挂载，可以实现文件的共享。