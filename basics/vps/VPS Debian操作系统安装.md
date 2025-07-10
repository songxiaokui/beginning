### 安装纯净的Debian OS
[参考地址](https://github.com/bohanyang/debi)

1. 下载脚本 
> curl -fLO https://raw.githubusercontent.com/bohanyang/debi/master/debi.sh && chmod a+rx debi.sh 

2. 安装
```shell
./debi.sh --bbr --cloud-kernel --ethx --timezone Asia/Shanghai --full-upgrade --version 12 --grub-timeout 0 --ssh-port 22 --user {你的用户名} --password {你的密码} --static-ipv4
```

3. 启动
```shell
sudo shutdown -r now
```