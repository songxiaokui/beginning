解决 utm 网络问题： 
    1. 关机 
    2. 先设置共享网络 
    3. 彻底释放旧的 IP 缓存 dhclient -r enp0s1 
    4. 重新向 UTM 申请新的 IP 地址 dhclient enp0s1