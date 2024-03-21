### 使用APDL导出网格信息  
- 通过ansys workbench创建一个分析模型  
- 导入几何  
- 设置几何材料  
- 设置边界条件(主要是用来设置网格中的组信息)  
- 在环境中->工具->生成MAPDL输入文件(.dat)
- 然后打开Ansys APDL软件
- 读入.dat文件
- 使用命令流导出  
  > cdwrite,db,'C:\\test',cdb