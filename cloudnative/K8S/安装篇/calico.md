### calico组件安装错误
1. 查询出现问题pod的具体信息
   kubectl describe pod calico-node-8bx4q -n kube-system
   错误信息如下:
   Events:
   Type     Reason     Age                  From               Message
   ----     ------     ----                 ----               -------
   Normal   Scheduled  3m10s                default-scheduler  Successfully assigned kube-system/calico-node-8bx4q to k8s-node02
   Normal   Pulling    3m9s                 kubelet            Pulling image "registry.cn-beijing.aliyuncs.com/dotbalo/cni:v3.26.1"
   Normal   Pulled     2m25s                kubelet            Successfully pulled image "registry.cn-beijing.aliyuncs.com/dotbalo/cni:v3.26.1" in 44.094s (44.094s including waiting)
   Normal   Created    47s (x5 over 2m25s)  kubelet            Created container upgrade-ipam
   Normal   Started    47s (x5 over 2m25s)  kubelet            Started container upgrade-ipam
   Normal   Pulled     47s (x4 over 2m24s)  kubelet            Container image "registry.cn-beijing.aliyuncs.com/dotbalo/cni:v3.26.1" already present on machine
   Warning  BackOff    47s (x9 over 2m23s)  kubelet            Back-off restarting failed container upgrade-ipam in pod calico-node-8bx4q_kube-system(687bbc84-69ac-4f31-95a1-8266c9d6f1e4)
2. 查询出现具体容器报错信息
   从上面可知是容器初始化报错 upgrade-ipam, 查询具体容器
   kubectl logs calico-node-8bx4q -c upgrade-ipam
   日志如下:
   exec /opt/cni/bin/calico-ipam: exec format error
   问题定位: 容器架构不匹配导致
3. 解决方案，去官方找calico.yaml使用arm架构容器包即可
4. 下载地址: https://projectcalico.docs.tigera.io/archive/v3.25/manifests/calico.yaml
5. 修改calico的cidr 使用pod网段即可，最好定义一个独立的网段