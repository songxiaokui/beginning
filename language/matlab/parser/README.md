### 环境搭建

```shell
python -m venv .matlab  

```

### 解析Matlab结果文件

```txt
parser/   
├── __init__.py  
├── loader/  
│   ├── __init__.py  
│   ├── base_loader.py        # 定义 Loader 接口  
│   ├── scipy_loader.py       # 适配 v5/v7  
│   └── hdf5_loader.py        # 适配 v7.3  
├── converter/  
│   ├── __init__.py  
│   ├── flatten_struct.py     # MATLAB 结构体展平为 dict  
│   └── export_csv.py         # 输出 CSV/JSON  
├── cli/  
│   ├── __init__.py  
│   └── main.py               # 命令行入口   
├── utils/  
│   ├── __init__.py  
│   └── log.py                # 日志输出工具  
├── requirements.txt  
└── README.md  
```

### 说明

什么是 MCOS？  
MCOS = MATLAB Class Object System
是 MATLAB 在 R2008a 引入的类机制，任何 datetime / string / table 等都属于 MATLAB 对象（内部是 classdef 定义的），存储在 .mat
文件中时会变成：`b'MCOS'`
在 HDF5 格式中，它是个对象引用数组（通常你看到 MATLAB_object, __properties__, __classname__ 字段）。

📦 .mat 文件的存储方式（按版本区分）：

| MATLAB 数据类型                                  | V5格式存储方式  | V7格式 | V7.3 (HDF5) 格式          |
|----------------------------------------------|-----------|------|-------------------------|
| `double`, `char`, `logical`                  | 原生支持      | 支持   | 原生存储                    |
| `datetime`, `string`, `categorical`, `table` | 以结构体或对象方式 | 不支持  | 以 MCOS 对象引用 存储          |
| 自定义 class / handle                           | 不支持       | 不支持  | MCOS 对象（必须 MATLAB 才能解析） |

注意:   
如果在保存数据序列化时，没有使用常见存储格式，而是使用对象格式，则Matlab会使用MCOS进行存储，此时数据通过Python是无法进行读取并解析，例如
datetime('now') 应该在序列化时  
存储为 datestr(datetime('now'), 'yyyy-mm-dd HH:MM:SS.FFF') 或者保存为时间戳：posixtime(datetime('now'))
，字符串也是如此，应该保存为char,  
例如：name = char('武安君') 而不要直接使用 name = "武安君"  
所以尽力使用扁平话结构  

说明:
1. main.py支持 v5、v7、v7.3 的带表头的数据格式
2. hdf5_loader.py 适配 v7.3 无表头的数据格式