#!/bin/bash

docker build ./ --tag istio-test/goods-http:v1

kubectl delete -f ./kubernetes/deployment.yml
kubectl apply -f ./kubernetes/deployment.yml
