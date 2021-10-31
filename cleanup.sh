#!/bin/sh

helm delete log-app
helm delete fluent-bit
kind delete cluster --name fluent-bit-test
rm -r /tmp/fluent-bit-test
