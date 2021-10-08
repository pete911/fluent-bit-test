# fluent-bit-test
project to manually test fluent-bit configurations

## requirements
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [helm](https://helm.sh/docs/helm/helm_install/) at least 3.7.0 version

## run
- modify [values.yml](values.yml) if needed and run `./test.sh` start kind cluster
- cluster runs test app exposed on `30500` port, which can be used to generate logs by POSTing data to it:
  - `curl -v -X POST -d 'test log message' http://localhost:30500`
- logs are sent to `/tmp/fluent-bit-test/` local directory
  - `tail -f /tmp/fluent-bit-test/kube.var.log.containers.log-app-*.log | jq .`
- clean up after test `kind delete cluster --name fluent-bit-test && rm -r /tmp/fluent-bit-test`

## debug kind cluster
- `docker ps` - list kind cluster nodes running as docker images
- `docker exec -it fluent-bit-test-control-plane bash` - "ssh" onto control plane node
- once on the node, we can:
    - `ps auxf` check processes
    - `crictl images` check images already present on the node
