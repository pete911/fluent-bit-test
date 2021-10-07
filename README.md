# fluent-bit-test
project to manually test fluent-bit configurations

## requirements
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [helm](https://helm.sh/docs/helm/helm_install/) at least 3.7.0 version

## run
- `./test.sh` start kind cluster
- logs are sent to `/tmp/fluent-bit-test/` directory
  - view fluent-bit logs `cat /tmp/fluent-bit-test/kube.var.log.containers.fluent-bit-*.log | jq .`
  - view etcd logs `cat /tmp/fluent-bit-test/kube.var.log.containers.etcd-*.log | jq .`
- `kind delete cluster --name fluent-bit-test && rm -r /tmp/fluent-bit-test` delete cluster and logs

## debug kind cluster
- `docker ps` - list kind cluster nodes running as docker images
- `docker exec -it fluent-bit-test-control-plane bash` - "ssh" onto control plane node
- once on the node, we can:
    - `ps auxf` check processes
    - `crictl images` check images already present on the node
