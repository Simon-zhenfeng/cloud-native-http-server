### 创建configmap
```shell
$ kubectl create configmap http-server-config --from-env-file=http-server.conf
```

### 创建deployment
```shell
$ kubectl create -f deployment.yaml
```

### 创建service
```shell
$ kubectl create -f service.yaml
```

### 安装nginx-ingress
```shell
sudo snap install helm --classic
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress
```

### 安装证书
```shell
helm repo add jetstack https://charts.jetstack.io
helm repo update
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.crds.yaml
helm install cert-manager jetstack/cert-manager --namespace cert-manager --create-namespace --version v1.8.0
  # --set installCRDs=true
```