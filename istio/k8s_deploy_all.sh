#!/bin/bash

kubectl apply -f ./user/kubernetes/deployment.yml
kubectl apply -f ./comments/kubernetes/deployment.yml
kubectl apply -f ./comments/kubernetes/deployment-v2.yml
kubectl apply -f ./stock/kubernetes/deployment.yml
kubectl apply -f ./goods/kubernetes/deployment.yml

kubectl apply -f ./user/kubernetes/service.yml
kubectl apply -f ./comments/kubernetes/service.yml
kubectl apply -f ./stock/kubernetes/service.yml
kubectl apply -f ./goods/kubernetes/service.yml
