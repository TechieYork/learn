#!/bin/bash

kubectl apply -f ./user-grpc/kubernetes/deployment.yml
kubectl apply -f ./comments-grpc/kubernetes/deployment.yml
kubectl apply -f ./comments-grpc/kubernetes/deployment-v2.yml
kubectl apply -f ./stock-grpc/kubernetes/deployment.yml
kubectl apply -f ./goods-http/kubernetes/deployment.yml
kubectl apply -f ./goods-grpc/kubernetes/deployment.yml

kubectl apply -f ./user-grpc/kubernetes/service.yml
kubectl apply -f ./comments-grpc/kubernetes/service.yml
kubectl apply -f ./stock-grpc/kubernetes/service.yml
kubectl apply -f ./goods-http/kubernetes/service.yml
kubectl apply -f ./goods-grpc/kubernetes/service.yml
