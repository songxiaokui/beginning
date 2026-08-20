# OpenFOAM Foundation 13

这个版本是 OpenFOAM Foundation 发行版，主要面向高校教学、科研和工程仿真使用。

本目录用于构建 OpenFOAM Foundation v13 的 Singularity/Apptainer 镜像。
这里的 `13` 指 OpenFOAM Foundation 的版本号，不是 OpenCFD 发行版的版本号；两者虽然都使用 OpenFOAM 名称，但代码分支、版本命名、求解器和字典格式并不完全兼容。

OpenFOAM Foundation v13 由 OpenFOAM Foundation 发布，适合高校教学、科研和需要从源码编译求解器的场景。官方版本信息见 [OpenFOAM 13](https://openfoam.org/version/13/) 和 [下载页面](https://openfoam.org/download/)。

## 目录内容

| 文件 | 作用 |
| --- | --- |
| `Makefile` | 下载源码并调用 SingularityCE 构建 arm64/amd64 镜像 |
| `openfoam13-multiarch.def` | Ubuntu 24.04、多架构源码编译定义文件 |
| `tar/` | OpenFOAM 13 与 ThirdParty 13 源码压缩包，构建时自动创建 |

## 与 OpenCFD 版本的区别

本目录不是 OpenCFD 的 `OpenFOAM-v2312`、`v2406`、`v2412` 或其他 `vYYMM` 版本。使用案例或自定义 solver 时，应先确认案例来自哪个发行版。

| 项目 | 本目录：Foundation | OpenCFD |
| --- | --- | --- |
| 版本命名 | `OpenFOAM-13`、`v13` | 通常为 `OpenFOAM-vYYMM` |
| 官方站点 | [openfoam.org](https://openfoam.org/) | [openfoam.com](https://www.openfoam.com/) |
| 发布节奏 | 以递增的大版本号发布 | 以年月版本号发布 |
| 兼容性 | 与 Foundation 分支的教程、案例和源码匹配 | 应使用对应 OpenCFD 版本的教程、案例和源码 |
| 本镜像的入口 | `openfoam13 <命令>` 或 `singularity run` | 不适用，不能仅替换镜像文件名继续使用 |

不要因为两个版本都存在 `simpleFoam`、`blockMesh` 等同名程序，就假定所有案例可以直接互换。`controlDict`、求解器选择、模型名称、字典路径和自定义库接口都可能存在差异；迁移案例时应以目标分支的官方文档和 `$FOAM_TUTORIALS` 为准。

## 构建环境和限制

- 基础镜像：Ubuntu 24.04。
- 构建方式：从源码编译 OpenFOAM Foundation 13。
- MPI：系统 OpenMPI，`WM_MPLIB=SYSTEMOPENMPI`。
- 支持架构：原生 `amd64` 和 `arm64`。
- 构建必须在与目标架构一致的主机/rootfs 上完成；当前定义文件不提供跨架构交叉编译。
- 默认并行度为 4，可通过 `build_jobs` 调整。
- 镜像内不编译 ParaView，构建时设置 `ParaView_TYPE=none`。

因此，在 ARM64 主机上构建 amd64，或在 amd64 主机上构建 arm64，会被定义文件主动拒绝。若需要另一种架构，应在对应架构的构建节点上执行构建。

## 构建镜像

先进入本目录：

```bash
cd tools/openfoam/13
```

### ARM64

```bash
make build-arm64
```

输出镜像：`openfoam-13-arm64.sif`。

### AMD64

```bash
make build-amd64
```

输出镜像：`openfoam-13-amd64.sif`。

### 手动构建或调整并行度

```bash
singularity build \
  --build-arg target_arch=arm64 \
  --build-arg build_jobs=8 \
  openfoam-13-arm64.sif \
  openfoam13-multiarch.def
```

`target_arch` 支持 `auto`、`amd64` 和 `arm64`。不传参数时使用 `auto`，由构建环境检测架构。源码包由 `make download` 下载到 `tar/`；如果压缩包已存在，Makefile 不会重复下载。

```bash
make download
make help
```

当前 Makefile 使用项目维护的源码归档地址；如归档地址或版本补丁发生变化，应同步更新 `Makefile` 中的文件名和 URL，并确认压缩包解压后目录名仍分别为 `OpenFOAM-13-version-13` 和 `ThirdParty-13-version-13`。

## 运行方式

### 推荐：使用 `run`

`%runscript` 会调用镜像内的 `openfoam13` 包装脚本，该脚本先加载 Foundation 13 的 `etc/bashrc`，再执行 OpenFOAM 命令：

```bash
singularity run openfoam-13-arm64.sif foamRun -help
singularity run openfoam-13-arm64.sif blockMesh -help
```

在案例目录中运行时，建议绑定案例目录并设置工作目录：

```bash
singularity run \
  -B "$(pwd):/case" \
  --pwd /case \
  /path/to/openfoam-13-arm64.sif \
  blockMesh
```

### 使用 `exec`

直接 `exec` 时应显式调用 `openfoam13`，以确保加载 OpenFOAM 环境：

```bash
singularity exec openfoam-13-arm64.sif openfoam13 foamRun -help
singularity exec openfoam-13-arm64.sif openfoam13 blockMesh
```

不要依赖下面这种未加载环境的调用方式：

```bash
singularity exec openfoam-13-arm64.sif foamRun -help
```

镜像的 environment section 只设置基础变量，OpenFOAM 的 PATH、库路径以及其他构建变量由 `OpenFOAM-13/etc/bashrc` 配置；`openfoam13` 是统一入口。

## 案例运行示例

```bash
cd /path/to/case

# 检查网格
singularity run \
  -B "$(pwd):/case" --pwd /case \
  /path/to/openfoam-13-arm64.sif \
  checkMesh

# 运行求解器；具体 solver 以案例 controlDict 和 Foundation v13 文档为准
singularity run \
  -B "$(pwd):/case" --pwd /case \
  /path/to/openfoam-13-arm64.sif \
  foamRun
```

如果在集群中使用 MPI，应让宿主机的 Slurm/资源分配和镜像内的 OpenMPI 配置保持一致。例如先申请资源，再在绑定的案例目录中执行 `mpirun` 或集群封装命令；不要把 OpenCFD 镜像中的 MPI 参数直接照搬到本镜像。

## 构建验证

定义文件在构建末尾会自动检查：

- `foamRun`、`blockMesh`、`decomposePar` 和 `reconstructPar` 是否存在。
- `foamRun` 是否可以启动。
- 动态库依赖中是否存在 `not found`。
- `WM_ARCH`、`WM_OPTIONS` 是否与实际架构匹配。
- MPI 是否为 `SYSTEMOPENMPI`。

构建完成后可再次执行：

```bash
singularity run openfoam-13-arm64.sif foamRun -help
singularity exec openfoam-13-arm64.sif openfoam13 bash -lc \
  'echo "WM_PROJECT_DIR=$WM_PROJECT_DIR"; echo "WM_OPTIONS=$WM_OPTIONS"; echo "WM_MPLIB=$WM_MPLIB"; command -v foamRun'
```

## 常见问题

### 找不到 `foamRun` 或 `blockMesh`

通常是直接使用了 `singularity exec image <command>`，没有经过环境初始化。改为：

```bash
singularity exec image openfoam13 foamRun -help
```

或直接使用 `singularity run image foamRun -help`。

### 架构不匹配

检查主机和容器构建环境：

```bash
uname -m
dpkg --print-architecture
```

二者应分别对应 `x86_64/amd64` 或 `aarch64/arm64`。这是源码原生编译限制，不是运行参数问题。

### OpenCFD 案例在 Foundation 13 中报错

先确认案例来源和版本，再按 Foundation v13 的文档调整字典、求解器和模型名称。不要只修改镜像名，也不要将 OpenCFD 的 `etc/bashrc`、环境变量或库文件复制到本镜像中。

### 如何查看完整编译日志

OpenFOAM 源码编译日志位于构建过程中的 `${WM_PROJECT_DIR}/log.Allwmake`。如果构建失败，应保留 Singularity 构建输出以及该日志中的第一个错误，优先检查依赖、源码包完整性和架构是否匹配。
