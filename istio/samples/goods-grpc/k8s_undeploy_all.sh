#!/bin/bash

kubectl delete -f ./user-grpc/kubernetes/deployment.yml
kubectl delete -f ./comments-grpc/kubernetes/deployment.yml
kubectl delete -f ./comments-grpc/kubernetes/deployment-v2.yml
kubectl delete -f ./stock-grpc/kubernetes/deployment.yml
kubectl delete -f ./goods-http/kubernetes/deployment.yml

kubectl delete -f ./user-grpc/kubernetes/service.yml
kubectl delete -f ./comments-grpc/kubernetes/service.yml
kubectl delete -f ./stock-grpc/kubernetes/service.yml
kubectl delete -f ./goods-http/kubernetes/service.yml
