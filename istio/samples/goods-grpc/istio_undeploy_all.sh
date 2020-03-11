#!/bin/bash

kubectl delete -f ./user-grpc/istio/virtual-service.yml
kubectl delete -f ./comments-grpc/istio/virtual-service.yml
kubectl delete -f ./stock-grpc/istio/virtual-service.yml
kubectl delete -f ./goods-http/istio/virtual-service.yml
kubectl delete -f ./goods-grpc/istio/virtual-service.yml

kubectl delete -f ./goods-http/istio/gateway.yml
kubectl delete -f ./goods-http/istio/gateway-virtual-service.yml

kubectl delete -f ./goods-grpc/istio/gateway.yml
kubectl delete -f ./goods-grpc/istio/gateway-virtual-service.yml

kubectl delete -f ./user-grpc/istio/destination-rule.yml
kubectl delete -f ./comments-grpc/istio/destination-rule.yml
kubectl delete -f ./stock-grpc/istio/destination-rule.yml
kubectl delete -f ./goods-http/istio/destination-rule.yml
kubectl delete -f ./goods-grpc/istio/destination-rule.yml
