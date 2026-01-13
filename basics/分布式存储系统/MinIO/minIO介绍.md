### MinIO介绍
> MinIO 是高性能、兼容 S3 的对象存储解决方案
> 
> 
### 文档
[文档地址](https://github.com/minio/minio) 


### 部署
* 单节点容器化部署  
```shell
docker rm -f minio-storage

mkdir -pv ./.rootfs

docker run -p 9001:9000 -p 9002:9002 \
  --name "minio-storage" \
  -itd \
  --restart=always \
  -e "MINIO_ROOT_USER=Admin" \
  -e "MINIO_ROOT_PASSWORD=Admin0987abc" \
  -v ./.rootfs:/data \
  quay.io/minio/minio server /data --console-address ":9002"

```

### 访问
> http://localhost:9002
> 

### 集成在k8s中(使用k3s模拟)
* Helm部署MinIO
  ```shell
  # 重新创建集群
  k3d cluster create test-cluster \
  --agents 1 \
  --agents-memory 8096m \
  --volume ~/minio-data:/mnt/minio-data@agent:0
  
  helm repo add minio https://charts.min.io/
  helm repo update
  kubectl create namespace minio
  
  # 由于我们mock，所以需要存储，就用pvc给minIO做存储，正常来说是minIO提供给其他Pod使用（云）
  mkdir -pv ~/minio/storage

  # 创建pvc
  kubectl apply -f minio-pvc.yaml
  
  # 安装minio
  helm install minio minio/minio --namespace minio -f minio-values.yaml

  # 验证pod状态
  kubectl get pods -n minio -o wide
  
  helm list -n minio
  kubectl get svc -n minio
  
  # 排错1：pod一直无法调度成功
  # 解决方案: 查看node的信息，是docker 存储不足
  
  # 排错2：pod启动，但是有一个容器启动失败
  # 解决方案：本地路径挂载权限不对，将pvc创建逻辑改成自动创建
  
  # 访问minio控制台 kubectl 端口转发
  # 出现无法访问，将类型改为NodePort模式 避免使用9000端口 sing-box默认使用了9000 所以改端口然后映射
  kubectl port-forward $(kubectl get pods -n minio  | grep minio | awk '{print $1}') 31002:9001 -n minio

  # 访问minio控制台
  http://127.0.0.1:31002/browser
  ```
* 安装csi-s3插件
* 创建MinIO的Secret
* 创建SC指向MinIO
* 创建PVC
* 创建Pod验证