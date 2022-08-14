[![License](https://img.shields.io/badge/license-Apache%202.0-brightgreen.svg)](License)
# nvd-data-mirror

为 maven 插件 org.owasp:dependency-check-maven 提供数据镜像服务，同时包含 cve 和 retireJs 数据。

## Usage

```bash
cd /path/to/nvd-data-mirror
# 查看配置说明
./nvd-data-mirror --help

# 启动服务，默认端口为 80
./nvd-data-mirror
```

## Docker

```bash
docker run -d -p 80:80 -v /path/to/nvd-data-mirror:/data nvd-data-mirror:latest
```

## Kubernetes


```bash
# helm 模式部署
helm repo add jing https://jingshouyan.github.io/helm-charts/
helm repo update

helm install nvd-data-mirror jing/nvd-data-mirror --set service.type=NodePort --set service.nodePort=30010

```

### helm chart 源码

https://github.com/jingshouyan/helm-charts/tree/main/charts/nvd-data-mirror

## 其他

参考 [nist-data-mirror](https://github.com/stevespringett/nist-data-mirror)
