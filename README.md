# Simple Kubernetes Operator

*This project only for personal use and study so it is very simple and not professional*.

*I'm happy if you contribute to this repo.*

* Server with grpc
* Get all namespaces
* Get all pods with specific namespace
* Get all services with specific namespace
* Get all deployments with specific namespace
* Delete specific pod

## Use

### Directly
```shell
# Your GRPC server port
# Default is 9090 if not set
export  GRPC_LISTEN_PORT=

# Kubernetes Master Url
# Optional
export K8S_MASTER_URL=

# Kubernetes Config File Path 
# Default is '/etc/kubernetes/admin.conf'
export K8S_CONFIG_FILE=

go build -o kube-operator

./kube-operator
```
### By docker

```shell
# Please make sure your Kubernetes configuration file is mounted in the specified directory. 
# The default directory can be found in [Directly format]
# And you can also customize it by adjusting the environment variables of the Docker container.
docker run --name msr_http kube-operator -d registry.cn-hangzhou.aliyuncs.com/lexmargin/kube-operator:latest

```

### By Kubernetes
```shell
# You may want to adjust the env value or node port by your self
kubectl apply -f https://github.com/NoahAmethyst/simple-kube-operator/blob/master/kube_operator.yml
```

