# OpenCFD OpenFOAM v2606

本目录用于构建 OpenCFD 发布的 OpenFOAM v2606 Singularity/Apptainer 镜像。它与 `tools/openfoam/13` 的 OpenFOAM Foundation v13 不属于同一条代码分支，不能直接混用两者的案例、环境脚本或自定义库。

官方资料：

- [v2606 发布说明](https://www.openfoam.com/news/main-news/openfoam-v2606)
- [v2606 源码仓库](https://gitlab.com/openfoam/core/openfoam/-/tree/OpenFOAM-v2606?ref_type=tags)
- [版本与源码下载](https://www.openfoam.com/download/release-history)

## 构建目标

| 目标 | 构建主机 | 输出 |
| --- | --- | --- |
| 麒麟 ARM64 | 麒麟 ARM64/aarch64 主机 | `openfoam-v2606-kylin-arm64.sif` |
| AMD64 | AMD64/x86_64 主机 | `openfoam-v2606-amd64.sif` |

SIF 中的程序与架构相关，请分别在两台对应架构的主机上构建。这里默认“麒麟”是 ARM64 构建主机，SIF 内部使用 Ubuntu 24.04 用户空间；如果需要使用麒麟作为 SIF 基础镜像，需要另行准备基础镜像和软件源。

## 构建

```bash
cd tools/openfoam/v2606
make build-kylin-arm64   # 麒麟 ARM64
make build-amd64         # AMD64
```

默认使用 4 个并行任务，并启用 OpenCFD 推荐的队列编译参数。根据构建机 CPU 和内存调整并行度：

```bash
make BUILD_JOBS=8 build-amd64
```

构建日志默认写入 `logs/`，标准输出和错误输出分开保存：

```text
logs/openfoam-v2606-amd64.stdout.log
logs/openfoam-v2606-amd64.stderr.log
logs/openfoam-v2606-kylin-arm64.stdout.log
logs/openfoam-v2606-kylin-arm64.stderr.log
```

可以通过 `LOG_DIR` 指定其他日志目录，例如：

```bash
make LOG_DIR=/data/openfoam-build-logs BUILD_JOBS=8 build-amd64
```

也可以手动构建：

```bash
singularity build \
  --build-arg target_arch=arm64 \
  --build-arg build_jobs=8 \
  openfoam-v2606-kylin-arm64.sif \
  openfoam-v2606-multiarch.def
```

仅下载源码：`make download`。源码文件为 `tar/OpenFOAM-v2606.tgz` 和 `tar/ThirdParty-v2606.tgz`，解压后必须分别生成 `/opt/openfoam/OpenFOAM-v2606` 和 `/opt/openfoam/ThirdParty-v2606`。

## 构建内容

- 基础镜像：Ubuntu 24.04；源码编译 OpenCFD v2606。
- 官方基础工具链：`build-essential`、`autoconf`、`autotools-dev`、`cmake`、`gawk`、`m4`、`flex`、`libfl-dev`。
- 官方数学/MPI 依赖：`libgmp-dev`、`libmpfr-dev`、`libmpc-dev`、OpenMPI 和 MPI 默认包装包。
- 官方系统数学/分区依赖：`libfftw3-dev`、`libscotch-dev`、`libptscotch-dev`、`libcgal-dev`。
- MPI：系统 OpenMPI，`WM_MPLIB=SYSTEMOPENMPI`。
- Flex：`flex` 与 `libfl-dev`，后者提供 `FlexLexer.h`。
- ParaView：不编译，设置 `ParaView_TYPE=none`。
- 入口：`/usr/local/bin/openfoam-v2606`。

OpenCFD 的 ThirdParty 构建默认还会编译 Boost、CGAL、FFTW、KaHIP 和 Scotch。除非同步修改 OpenFOAM 配置改用系统库，否则单独安装 `libfftw3-dev`、`libscotch-dev`、`libptscotch-dev` 或 `libcgal-dev` 不会自动跳过这些源码构建；因此当前 definition 保持官方默认的 ThirdParty 路径。

关键路径为 `WM_PROJECT_DIR=/opt/openfoam/OpenFOAM-v2606` 和 `WM_THIRD_PARTY_DIR=/opt/openfoam/ThirdParty-v2606`。

## 运行

推荐使用 `run`：

```bash
singularity run openfoam-v2606-kylin-arm64.sif simpleFoam -help
singularity run openfoam-v2606-kylin-arm64.sif blockMesh -help
```

镜像入口已经使用独立的 `bash --noprofile --norc` 加载 OpenCFD 环境，不会读取宿主机的 Bash 初始化文件。若宿主机使用环境模块或存在复杂的容器环境注入，仍可显式使用 `--cleanenv`：

```bash
singularity run --cleanenv openfoam-v2606-kylin-arm64.sif blockMesh -help
```

使用 `exec` 时显式调用统一入口：

```bash
singularity exec openfoam-v2606-kylin-arm64.sif openfoam-v2606 simpleFoam -help
```

案例目录建议绑定并设置工作目录：

```bash
cd /path/to/case
singularity run -B "$(pwd):/case" --pwd /case \
  /path/to/openfoam-v2606-amd64.sif blockMesh
```

不要直接使用 `singularity exec image.sif simpleFoam -help`，因为它不会保证 OpenCFD 环境已经加载。

## 验证和故障排查

构建阶段执行官方的 `foamSystemCheck`、`Allwmake` 和 `foamInstallationTest`。构建日志位于镜像内 `/opt/openfoam/OpenFOAM-v2606/log.linux64GccDPInt32Opt`。

主机检查：

```bash
uname -m
dpkg --print-architecture
```

麒麟 ARM64 应对应 `aarch64`/`arm64`，AMD64 应对应 `x86_64`/`amd64`。OpenCFD v2606 与 Foundation v13 即使存在同名程序，也不代表案例、字典、模型名称和自定义库可以直接互换。
