### 创建configmap
```shell
$ kubectl create configmap http-server-config --from-env-file=http-server.conf
```

### 创建deployment
```shell
$ kubectl create -f deployment.yaml
```