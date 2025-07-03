### 安装Matlab软件R2021a+

[Window10安装](https://www.cnblogs.com/sixuwuxian/p/13973003.html)

[官方地址](https://www.mathworks.com/)

[Macbook安装](https://zhuanlan.zhihu.com/p/681631044)

Matlab是需要License的，本地测试只是想用Matlab生成一些test数据，暂时按照破解版，支持多个版本的数据

推荐安装: R2021a+


## 版本选择建议（兼容 `.mat` V5 / V7 / V7.3）

| MATLAB `.mat` 格式 | 默认支持版本                             | 安装建议          |
| ------------------ | ---------------------------------------- | --------------- |
| V5/V7              | 所有版本（R2006b 及以上）                | 支持            |
| V7.3               | R2006b 及以上，但需开启 `-v7.3` 存储选项 | 支持（基于 HDF5） |


说明: Windows 7 下安装 MATLAB. **MATLAB R2019b 是官方最后支持 Windows 7 的版本**，R2020a 起官方不再支持 Win7,所以推测现场数据是不支持HDF5存储

### 安装步骤

1. 下载   
   https://link.zhihu.com/?target=https%3A//pan.baidu.com/s/12wArirw5Pk_nHysjbH5Tvg
   提取码: 2t9z

2. 安装  
   关闭代理VPN
   密钥：62551-02011-26857-57509-64399-54230-13279-37181-62117-65158-40352-64197-45508-24369-45954-39446-39538-16936-10698-58393-44718-32560-10501-40058-34454
   安装目录: /Applications/matlab
   License: crack/license.lic

3. 替换  
   cp crack/libmwlmgrimpl.dylib /Applications/matlab/MATLAB_R2021b.app/bin/maci64/matlab_startup_plugins/lmgrimpl/libmwlmgrimpl.dylib
4. matlab 一个文件一个函数，函数名等于文件名