![build](https://img.shields.io/github/workflow/status/jingshouyan/nvd-data-mirror/Release%20Go%20project)
![release](https://img.shields.io/github/v/release/jingshouyan/nvd-data-mirror)
![version](https://img.shields.io/github/go-mod/go-version/jingshouyan/nvd-data-mirror)
[![License](https://img.shields.io/badge/license-Apache%202.0-brightgreen.svg)](LICENSE)
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

maven 配置

${nvd-data-mirror-host} 替换为服务器地址

```xml
    <plugin>
        <groupId>org.owasp</groupId>
        <artifactId>dependency-check-maven</artifactId>
        <version>${dependency-check-maven.version}</version>
        <executions>
            <execution>
                <goals>
                    <goal>check</goal>
                </goals>
                <configuration>
                    <cveUrlBase>${nvd-data-mirror-host}/data/nvdcve-1.1-%d.json.gz</cveUrlBase>
                    <cveUrlModified>${nvd-data-mirror-host}/data/nvdcve-1.1-modified.json.gz</cveUrlModified>
                    <retireJsUrl>${nvd-data-mirror-host}/data/jsrepository.json</retireJsUrl>
                </configuration>
            </execution>
        </executions>
    </plugin>
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
