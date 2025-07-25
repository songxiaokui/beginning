### K3D创建K8S集群

* 安装k3d
  ```shell
  # mac os
  brew install k3d
  # 查看是否可升级版本 brew outdated k3d
  # brew upgrade k3d
  # 查看版本号
  k3d version
  
  # linux install
  curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
  ```

* 创建集群
  ```shell
  k3d cluster create llm-cluster --servers 3 --agents 3 --api-port 127.0.0.1:6445 --k3s-arg "--disable=traefik@server:0" -p "9080:31080@loadbalancer" -p "9443:31443@loadbalancer"
  ```

* 导出集群配置
  ```shell
  k3d kubeconfig get llm-cluster > llm_kubeconfig.yaml
  ```

* helm使用集群
  ```shell
  export KUBECONFIG=llm_kubeconfig.yaml
  kubectl get nodes
  # 创建ns
  kubectl create ns llm
  # 部署nginx
  helm repo add bitnami https://charts.bitnami.com/bitnami
  helm repo update
  helm install llm-nginx bitnami/nginx --namespace llm --set replicaCount=3
  # 端口转发
  kubectl port-forward svc/llm-nginx 8081:80 -n llm
  curl http://127.0.0.1:8081
  
  # 安装nginx-ingress 控制器
  helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
  helm repo update
  # helm install ingress-nginx ingress-nginx/ingress-nginx --namespace ingress-nginx --set controller.service.type=LoadBalancer
  helm upgrade --install ingress-nginx ingress-nginx/ingress-nginx \
  --namespace ingress-nginx --create-namespace \
  --set controller.service.type=NodePort \
  --set controller.service.nodePorts.http=31080 \
  --set controller.service.nodePorts.https=31443
  
  kubectl get svc -n ingress-nginx
  
  # 配置ingress 规则
  kubectl apply -f nginx-ingress.yaml
  kubectl get ingress -n llm
  
  # 配置本地hosts映射
  vim /etc/hosts
  # 增加下面内容
  127.0.0.1 nginx.austsxk.com
  
  # 抓ingress访问日志
  kubectl logs -n ingress-nginx -l app.kubernetes.io/component=controller --tail=50 -f
  
  # 访问服务
  curl -v http://nginx.austsxk.com:9080
  ```

* 删除集群
  ```shell
  k3d cluster delete llm-cluster
  ```

