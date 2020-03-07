#!/bin/bash

kubectl delete -f ./user/istio/virtual-service.yml
kubectl delete -f ./comments/istio/virtual-service.yml
kubectl delete -f ./stock/istio/virtual-service.yml
kubectl delete -f ./goods/istio/virtual-service.yml

kubectl delete -f ./goods/istio/gateway.yml

kubectl delete -f ./user/istio/destination-rule.yml
kubectl delete -f ./comments/istio/destination-rule.yml
kubectl delete -f ./stock/istio/destination-rule.yml
kubectl delete -f ./goods/istio/destination-rule.yml
