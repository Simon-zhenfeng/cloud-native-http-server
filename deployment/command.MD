### 创建configmap
```shell
$ kubectl create configmap http-server-config --from-file=config.yaml
```

### 创建deployment
```shell
$ kubectl create -f deployment.yaml
```