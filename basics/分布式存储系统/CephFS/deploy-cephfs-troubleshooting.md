# ✅ CephFS Docker 单节点部署错误排查全记录（ARM 架构）

## 📁 目录结构

```
ceph-cluster/
├── cephfs-server/
│   ├── .env                  # 含 FSID，务必与 ceph.conf 一致
│   ├── ceph.conf             # 包含 mon_host, fsid, auth_supported = none
│   ├── docker-compose.arm64.yml
│   ├── init-cephfs.sh        # 初始化 CephFS 脚本
│   └── mon-entrypoint.sh     # mon 节点初始化脚本
```

---

## 🚨 部署过程中遇到的错误与解决方案

### ❌ 1. `--gen-key: command not found`
**原因**：Compose 中多行 `command` 使用 `\` 续行语法导致 shell 解释错误。

**解决**：
- 改为单行：
  ```bash
  ceph-authtool --create-keyring ... --gen-key ...
  ```

---

### ❌ 2. `/init.sh: Is a directory`
**原因**：Makefile 中写了 `-v $(pwd)/init-cephfs.sh:/init.sh`，但 `$(pwd)` 在 Makefile 中无效。

**解决**：
- 使用 `$(shell pwd)` 替代：
  ```makefile
  -v $(shell pwd)/cephfs-server/init-cephfs.sh:/init.sh
  ```

---

### ❌ 3. `ceph-mon` 启动后容器自动退出（Exited 0）
**原因**：`ceph-mon` 启动在后台，容器没前台进程保持存活。

**解决**：
- 加 `--foreground` 参数：
  ```bash
  exec ceph-mon -i $(hostname -s) --keyring ... --foreground
  ```

---

### ❌ 4. `ceph auth get-or-create ...: No such file or directory`
**原因**：尝试写入 keyring 文件的目录不存在。

**解决**：
- 在 `init-cephfs.sh` 中增加：
  ```bash
  mkdir -p /var/lib/ceph/mgr/ceph-mgr
  ```

---

### ❌ 5. `mon_cmd_maybe_osd_create fail: 'wrong fsid'`
**原因**：OSD 数据目录残留旧的 FSID，但你 `.env` 中使用了新 FSID。

**解决**：
```bash
docker compose -f cephfs-server/docker-compose.arm64.yml down -v
rm -rf cephfs-server/data cephfs-server/osd
uuidgen > tmp && sed 's/.*/FSID=&/' tmp > cephfs-server/.env && rm tmp
```
- 替换 `ceph.conf` 中 `fsid = ...`，保持与 `.env` 一致

---

### ❌ 6. `no active mgr`、`mgr` 容器重启
**原因**：
- `mgr.mgr` 没有生成 keyring
- `ceph-mgr` 命令未加 `--foreground`

**解决**：
- 在 `init-cephfs.sh` 中创建 keyring
- Compose 中使用：
  ```yaml
  command: ceph-mgr -i mgr --foreground
  ```

---

### ❌ 7. `mds` 容器不断重启
**原因**：未成功创建 CephFS，或未生成 `mds` 的 keyring

**解决**：
- 确保 `init-cephfs.sh` 中执行：
  ```bash
  ceph fs new ...
  ceph auth get-or-create mds.mds-a ...
  ```

---

## ✅ 推荐部署流程

```bash
make config         # 自动生成 .env 和 ceph.conf（建议添加）
make up             # 启动集群服务
make init           # 初始化 pools、CephFS、keyring
make restart        # 重启 mgr、osd、mds 等组件
```

验证状态：

```bash
docker exec -it ceph-mon ceph -s
```

应输出：

```
mon: 1 daemons, quorum master
mgr: active
osd: 1 up, 1 in
mds: 1 up:active
```