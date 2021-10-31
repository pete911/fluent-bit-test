# fluent-bit-test
Project to manually test fluent-bit configurations. Includes an example of fluent-bit output plugin written in go.

## requirements
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
- [helm](https://helm.sh/docs/helm/helm_install/) at least 3.7.0 version
- [go](https://golang.org)

## run
- modify [values.yml](values.yml) if needed and run `make run` start kind cluster
- cluster runs test [log-app](log-app) exposed on `30500` port, which can be used to generate logs by POSTing data to it:
  - `curl -v -X POST -d 'test log message' http://localhost:30500`
- logs are sent to `/tmp/fluent-bit-test/` local directory
  - `tail -f /tmp/fluent-bit-test/kube.var.log.containers.log-app-*.log | jq .`
- project is also configured with [out-plugin](out-plugin) written in go. To view the output from this plugin, run:
  - `tail -f /tmp/fluent-bit-test/out`
- clean up after test `make cleanup`

## debug kind cluster
- `docker ps` - list kind cluster nodes running as docker images
- `docker exec -it fluent-bit-test-control-plane bash` - "ssh" onto control plane node
- once on the node, we can:
    - `ps auxf` check processes
    - `crictl images` check images already present on the node
