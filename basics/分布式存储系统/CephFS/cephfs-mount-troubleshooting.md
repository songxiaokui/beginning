# 📦 CephFS 挂载问题记录与解决方案

---

## ✅ 挂载最终成功说明

当你在客户端执行以下命令后成功：

```bash
sudo mount -t ceph 172.16.56.4:6789:/ /mnt/cephfs -o name=admin,secretfile=/etc/ceph/real.secret
```

代表你的 CephFS 文件系统部署和挂载全部完成。

---

## 🧨 客户端常见挂载失败问题汇总

### ❌ 报错：`secret is not valid base64`

**原因：**
传入的 `secretfile` 内容不是合法 base64 格式。

**解决：**
使用 Ceph 工具生成合法密钥或使用伪造合法密钥：

```bash
echo "AQBXZTZkqLYaBxAAz4JtHtW8DzDiRjX/7FSbpg==" | sudo tee /etc/ceph/real.secret > /dev/null
sudo chmod 600 /etc/ceph/real.secret
```

---

### ❌ 报错：`mount option secret requires a value`

**原因：**
传了 `secret=` 但没提供合法值。

**解决：**
不能使用空密钥，必须是 base64 格式的字符串。

---

### ❌ 报错：`no mds server is up or the cluster is laggy`

**原因：**
服务端虽然有 mds 容器，但底层 pool 的 PG 状态异常，CephFS 不可用。

**解决：**
- 设置 pool 副本数为 1（本地部署）：
  ```bash
  ceph config set global mon_allow_pool_size_one true
  ceph osd pool set cephfs_data size 1 --yes-i-really-mean-it
  ceph osd pool set cephfs_metadata size 1 --yes-i-really-mean-it
  ```

---

## 🧨 服务端问题记录

### ❌ MDS 容器不断重启

**错误日志：**

```
*** got signal Terminated ***
```

**原因：**
- 容器中挂载的 `keyring` 缺失
- 宿主机上未提前写入 keyring 文件

**解决：**

```bash
mkdir -p cephfs-server/data/mds/ceph-mds-a
docker exec ceph-mon ceph auth get-or-create mds.mds-a mon 'allow profile mds' osd 'allow *' mds 'allow'   > cephfs-server/data/mds/ceph-mds-a/keyring
docker restart ceph-mds
```

---

### ❌ 设置 pool 副本数失败

**错误：**
```
Error EPERM: configuring pool size as 1 is disabled by default.
```

**解决方案：**

```bash
ceph config set global mon_allow_pool_size_one true
ceph osd pool set cephfs_data size 1 --yes-i-really-mean-it
```

---

## ✅ 挂载验证指令

```bash
mount | grep ceph
ls -l /mnt/cephfs
echo "hello from client" > /mnt/cephfs/test.txt
```

---

## 🎯 推荐挂载脚本封装

```bash
#!/bin/bash
MON_IP=${1:-"172.16.56.4"}
sudo mkdir -p /mnt/cephfs
sudo mount -t ceph ${MON_IP}:6789:/ /mnt/cephfs -o name=admin,secretfile=/etc/ceph/real.secret
```

---

> 适用于本地 ARM 架构下使用 Docker Compose 部署的 CephFS 单节点开发集群。