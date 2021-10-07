#!/bin/sh

dir=$(dirname "$0")
mkdir -p /tmp/fluent-bit-test

echo " --- [start cluster] ---"
kind create cluster --config "$dir"/kind-config.yml --wait 5m --name fluent-bit-test

if [ "$(kubectl config current-context)" != "kind-fluent-bit-test" ]
then
  echo "kubectl context is not set to kind-fluent-bit-test"
  exit 1
fi

# install fluent-bit
helm repo add fluent https://fluent.github.io/helm-charts
helm repo update fluent
helm install --values "$dir"/values.yml fluent-bit fluent/fluent-bit

echo " --- [cluster running] ---"
