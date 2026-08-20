# OpenFoam 求解器

## OpenFoam-Foundation-V13

| 系统架构            | 镜像地址                                                     |
| ------------------- | ------------------------------------------------------------ |
| X86_64              | http://192.168.5.190:9444/images/openfoam-foundation-13-amd64.sif |
| Arrch64/arm64/Kylin | http://192.168.5.190:9444/images/openfoam-foundation-13-arm64.sif |

## OpenFoam-OpenCFD-V2606

| 系统架构            | 镜像地址                                                     |
| ------------------- | ------------------------------------------------------------ |
| X86_64              | http://192.168.5.190:9444/images/openfoam-OpenCFD-v2606-amd64.sif |
| Arrch64/arm64/Kylin | http://192.168.5.190:9444/images/openfoam-OpenCFD-v2606-arm64.sif |

### 案例包地址

http://192.168.5.190:9444/images/testcase-v2606.tar

## 镜像测试

需要在具备 Singularity 的环境下执行：

```bash
tar -zxvf testcase-v2606.tar

cd testcase-v2606/ && ./__script.sh ../openfoam-OpenCFD-v2606-amd64.sif
```

![openfoam-test](https://files.austsxk.com/2026-07-15/000001_openfoam.png)

