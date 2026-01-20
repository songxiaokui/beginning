# Docker buildx 磁盘清理与使用规范

> 适用场景：开发机长期使用 `docker buildx` 构建镜像，磁盘占用异常增长（数百 GB）。
>
> 目标：**安全清理历史 buildx 缓存 + 给出后续不复发的使用规范**。

---

## 一、问题背景

在开发机上长期使用 `docker buildx` 进行镜像构建（尤其是 `--push`、多次 `buildx create --use`），会产生大量 **BuildKit 状态缓存**：

* `buildx_buildkit_<name>0` 容器（buildkitd）
* `buildx_buildkit_<name>0_state` volume（真正占磁盘）

这些 **state volume 不会自动回收**，即使 builder 已经停止或废弃，也会长期占用磁盘，单个 volume 可达数十 GB。

---

## 二、根因总结

1. 多次执行 `docker buildx create --use`，生成大量 builder
2. 每个 builder 都会绑定一个持久化 `*_state` volume
3. `docker builder prune -a` **不会删除 state volume**
4. 长期未手动 `rm builder / rm volume`，导致磁盘累计爆满

---

## 三、排查步骤

### 1. 查看所有 buildx builder

```bash
docker buildx ls
```

判断规则：

* 带 `*`：当前正在使用的 builder（**不要删**）
* `inactive / stopped`：历史遗留 builder（**可删**）

---

### 2. 查看 buildkit 容器与 state volume

```bash
docker system df -v | grep buildx_buildkit
```

示例：

```text
buildx_buildkit_silly_keldysh0              Up
buildx_buildkit_silly_keldysh0_state        USED=1

buildx_buildkit_sad_khayyam0                Exited
buildx_buildkit_sad_khayyam0_state          USED=0   89.93GB
```

判断规则：

* `Up` + `USED=1`：当前 builder（保留）
* `Exited` + `USED=0`：历史垃圾缓存（可删）

---

## 四、安全清理流程（已验证）

> **严格顺序：先删容器 → 再删 volume**

---

### Step 1：删除已停止的 buildkit 容器

仅删除状态为 `Exited` 的容器：

```bash
docker rm buildx_buildkit_sad_khayyam0
docker rm buildx_buildkit_awesome_vaughan0
docker rm buildx_buildkit_multiarch0
docker rm buildx_buildkit_angry_einstein0
```

说明：

* 不影响任何已构建镜像或运行中的容器
* 仅用于解除 volume 的占用关系

---

### Step 2：删除真正占磁盘的 state volume

确认条件：

* `docker system df -v` 中显示 `USED=0`

逐个安全删除：

```bash
docker volume rm buildx_buildkit_nifty_darwin0_state
docker volume rm buildx_buildkit_sad_khayyam0_state
docker volume rm buildx_buildkit_intelligent_wilson0_state
docker volume rm buildx_buildkit_angry_einstein0_state
docker volume rm buildx_buildkit_multiarch0_state
docker volume rm buildx_buildkit_interesting_wozniak0_state
docker volume rm buildx_buildkit_infallible_cray0_state
docker volume rm buildx_buildkit_awesome_vaughan0_state
```

效果：

* 立即释放 **100GB+ 磁盘空间**
* 不影响当前构建环境

---

### 明确保留项（不要删除）

```text
buildx_buildkit_silly_keldysh0
buildx_buildkit_silly_keldysh0_state
```

原因：

* 当前 active builder（`Up` 状态）
* 后续构建仍在使用

---

## 五、清理验证

```bash
docker system df -v | grep buildx_buildkit
```

期望结果：

```text
buildx_buildkit_silly_keldysh0
buildx_buildkit_silly_keldysh0_state
```

---

## 六、后续构建规范（防止复发）

### 1. 开发机构建（推荐）

```bash
docker buildx build --no-cache --push ...
```

原因：

* 避免本地 buildkit cache 累积
* 构建结果直接推送仓库，不依赖本地缓存

---

### 2. CI / 发布环境（进阶）

```bash
docker buildx build \
  --cache-to=type=registry,ref=<image>:buildcache,mode=max \
  --cache-from=type=registry,ref=<image>:buildcache \
  --push ...
```

* cache 存储在 registry
* 本地磁盘可控

---

### 3. 定期维护（可选）

```bash
docker builder prune -a
docker volume prune
```

---

## 七、经验总结

> **buildx 的 state volume 不会自动清理，
> 不删除 builder 或 volume，磁盘一定会慢慢被吃满。**

---

## 八、适用范围

* Docker Desktop（macOS / Linux）
* 使用 buildx + BuildKit 的开发机
* 多语言 / 多阶段 / 多平台镜像构建场景

---

（完）
