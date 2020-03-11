#!/bin/bash

kubectl apply -f ./user-grpc/istio/virtual-service.yml
kubectl apply -f ./comments-grpc/istio/virtual-service.yml
kubectl apply -f ./stock-grpc/istio/virtual-service.yml
kubectl apply -f ./goods-http/istio/virtual-service.yml
kubectl apply -f ./goods-grpc/istio/virtual-service.yml

#kubectl apply -f ./goods-http/istio/gateway.yml
#kubectl apply -f ./goods-http/istio/gateway-virtual-service.yml

kubectl apply -f ./goods-grpc/istio/gateway.yml
kubectl apply -f ./goods-grpc/istio/gateway-virtual-service.yml

kubectl apply -f ./user-grpc/istio/destination-rule.yml
kubectl apply -f ./comments-grpc/istio/destination-rule.yml
kubectl apply -f ./stock-grpc/istio/destination-rule.yml
kubectl apply -f ./goods-http/istio/destination-rule.yml
kubectl apply -f ./goods-grpc/istio/destination-rule.yml
