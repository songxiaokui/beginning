### Windows安装OpenSSH-server

#### Windows Server
1. 安装
```shell
# 以管理员打开PowerShell
# 查询
Get-WindowsCapability -Online | Where-Object Name -like 'OpenSSH*'
# 如果 Server 是 NotPresent 安装
Add-WindowsCapability -Online -Name OpenSSH.Server~~~~0.0.1.0
```
2. 启动
```shell
# 以管理员打开PowerShell
Start-Service sshd
# 开机自启动
Set-Service -Name sshd -StartupType Automatic
```
3. 防火墙设置
```shell
New-NetFirewallRule -Name sshd -DisplayName "OpenSSH Server" -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 22
```
4. IP查询
```shell
ipconfig
```
5. 配置电脑不睡眠
```shell
# 永不睡眠
powercfg /change standby-timeout-ac 0
powercfg /change standby-timeout-dc 0

# 屏幕关闭时间
powercfg /change monitor-timeout-ac 10

# 关闭休眠
powercfg -h off
```
---
#### Client(mac)
1. 免密登陆
```shell
# ssh-copy-id -i ~/.ssh/xxxx.pub xxxx@IP
# 不起作用
# 执行
# 切换PowerShell
pwsh
whoami /groups | findstr /i administrators
# 创建文件
New-Item -ItemType File -Path "C:\ProgramData\ssh\administrators_authorized_keys" -Force
# 修复权限
icacls "C:\ProgramData\ssh\administrators_authorized_keys" /inheritance:r
icacls "C:\ProgramData\ssh\administrators_authorized_keys" /grant "Administrators:F"
icacls "C:\ProgramData\ssh\administrators_authorized_keys" /grant "SYSTEM:F"
# 启用零时编辑
icacls "C:\ProgramData\ssh\administrators_authorized_keys" /grant "你的用户名字:F"
# 内容写入
# 将mac上的~/.ssh/*.pub 的内容贴进去
# 重启sshd
Restart-Service sshd
```

