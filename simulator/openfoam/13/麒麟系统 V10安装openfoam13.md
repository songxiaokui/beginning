## [源码编译安装](https://openfoam.org/download/13-source/)

> 麒麟系统运行在国产鲲鹏、飞腾等 ARM 架构处理器上，或者方案一因版本依赖冲突无法使用，建议直接编译源码
>

### 1. 安装编译环境与依赖(容器化)
> 本次操作都在 arm 下的 singularity-v4.1.5 下进行
>
```shell
singularity build --sandbox openfoam-sd docker://ubuntu:24.04
singularity shell --writable openfoam-sd/

# 沙盒构建流程 root
apt-get update

apt-get clean

rm -rf /var/cache/apt/archives/partial
mkdir -p /var/cache/apt/archives/partial

# apt 下载直接使用 root 用户
apt-get \
    -o APT::Sandbox::User=root \
    install -y \
    build-essential \
    flex \
    bison \
    git \
    wget \
    cmake \
    zlib1g-dev \
    libboost-system-dev \
    libboost-thread-dev \
    libopenmpi-dev \
    openmpi-bin \
    gnuplot \
    libreadline-dev \
    libncurses-dev \
    libxt-dev
    
# 下载 openfoam
export OPENFOAM_PATH="/opt/openfoam"
mkdir -pv ${OPENFOAM_PATH}
cd ${OPENFOAM_PATH}

wget -O - http://dl.openfoam.org/source/13 | tar xvz
wget -O - http://dl.openfoam.org/third-party/13 | tar xvz
mv OpenFOAM-13-version-13 OpenFOAM-13
mv ThirdParty-13-version-13 ThirdParty-13

# 环境变量配置
export WM_PROJECT_DIR="${OPENFOAM_PATH}/OpenFOAM-13"
source ${WM_PROJECT_DIR}/etc/bashrc
echo $WM_PROJECT_DIR
echo $WM_THIRD_PARTY_DIR
echo $WM_ARCH
echo $WM_OPTIONS
echo $WM_MPLIB
echo $FOAM_MPI

# output
# /opt/openfoam/OpenFOAM-13
# /opt/openfoam/ThirdParty-13
# linuxArm64
# linuxArm64GccDPInt32Opt
# SYSTEMOPENMPI
# openmpi-system

# 禁用 export ParaView_TYPE=system
# OpenFOAM 官方源码明确支持从这些位置加载 prefs.sh 官方建议通过 prefs.sh 覆盖默认配置，而不是直接修改基础环境配置
# ~/.OpenFOAM/13
# ~/.OpenFOAM
# $FOAM_INST_DIR/site/13
# $FOAM_INST_DIR/site

mkdir -p /root/.OpenFOAM
cat > /root/.OpenFOAM/prefs.sh <<'EOF'
export ParaView_TYPE=none
EOF

source ${WM_PROJECT_DIR}/etc/bashrc
echo $ParaView_TYPE

# ARM64 编译 OpenFOAM 13
cd ${WM_PROJECT_DIR}
source ${WM_PROJECT_DIR}/etc/bashrc
# 预计要编译 8h
./Allwmake -j 8 2>&1 | tee log.Allwmake

# 编译验证
source ${WM_PROJECT_DIR}/etc/bashrc
# 查询
which foamRun
which blockMesh
which decomposePar
which reconstructPar
# 验证架构
file "$(which foamRun)"
foamRun -help

# 跑官方串行案例
find "$FOAM_TUTORIALS" \
    -type f \
    -name Allrun \
    | grep cavity \
    | head

td=`find "$FOAM_TUTORIALS" \
    -type f \
    -name Allrun \
    | grep cavity \
    | head`

echo $td
# /opt/openfoam/OpenFOAM-13/tutorials/legacy/incompressible/icoFoam/cavity/Allrun
mkdir -pv /opt/openfoam-test
cd /opt/openfoam-test
cp -r  $(dirname "${td}") .
cd cavity
./Allclean
./Allrun  2>&1 | tee arm-log.foamRun
echo $?
rm -rf /opt/openfoam-test

# 单节点 OpenMPI 4 Rank 并行验证
source ${WM_PROJECT_DIR}/etc/bashrc
mkdir -pv /opt/openfoam-parallel-test
cd /opt/openfoam-parallel-test
cp -r  $(dirname "${td}") .
cd cavity
./Allclean
cd cavity
ls -hl 
# 0  constant  system
# 生成网格
blockMesh 2>&1 | tee log.blockMesh
tail -5 log.blockMesh | grep "End" | wc -l
# 1
# 创建 decomposeParDict 使用scotch
cat > system/decomposeParDict <<'EOF'
FoamFile
{
    format      ascii;
    class       dictionary;
    object      decomposeParDict;
}

numberOfSubdomains 4;

method          scotch;
EOF
# 执行域分解
rm -rf processor*
decomposePar 2>&1 | tee log.decomposePar
tail -5 log.decomposePar | grep "End" | wc -l
ls -d processor*
# processor0  processor1	processor2  processor3
# 测试 mpi
mpirun --allow-run-as-root \
    -np 4 \
    bash -c \
    'echo rank=${OMPI_COMM_WORLD_RANK} size=${OMPI_COMM_WORLD_SIZE} host=$(hostname)'
  
mpirun --version
mpirun \
    --allow-run-as-root \
    -np 4 \
    icoFoam -parallel \
    2>&1 | tee log.icoFoam.parallel
# 重构结果
reconstructPar \
    2>&1 | tee log.reconstructPar
foamListTimes
# 输出
# 0.1
# 0.2
# 0.3
# 0.4
# 0.5
# 清理
rm -rf /opt/openfoam-parallel-test

# 增加统一入口
cat > /usr/local/bin/openfoam13 <<'EOF'
#!/bin/bash

source /opt/openfoam/OpenFOAM-13/etc/bashrc || exit 1

exec "$@"
EOF
chmod 755 /usr/local/bin/openfoam13
openfoam13 foamRun -help

# 如果用 def 记得source /opt/openfoam/OpenFOAM-13/etc/bashrc 写入 env 文件
cat >> /.singularity.d/env/90-environment.sh < 'EOF'
export LC_ALL="C.UTF-8"
export LANG="C.UTF-8"
. /opt/openfoam/OpenFOAM-13/etc/bashrc || exit 1
EOF


#### 退出容器
# 验证
singularity run openfoam-sd openfoam13 foamRun -help

# 构建镜像
singularity build openfoam_13.sif openfoam-sd/

# 测试
singularity exec  openfoam_13.sif openfoam13 foamRun -help
```