#!/bin/bash

kubectl apply -f ./user/istio/virtual-service.yml
kubectl apply -f ./comments/istio/virtual-service.yml
kubectl apply -f ./stock/istio/virtual-service.yml
kubectl apply -f ./goods/istio/virtual-service.yml

kubectl apply -f ./goods/istio/gateway.yml

kubectl apply -f ./user/istio/destination-rule.yml
kubectl apply -f ./comments/istio/destination-rule.yml
kubectl apply -f ./stock/istio/destination-rule.yml
kubectl apply -f ./goods/istio/destination-rule.yml
