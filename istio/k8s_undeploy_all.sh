#!/bin/bash

kubectl delete -f ./user/kubernetes/deployment.yml
kubectl delete -f ./comments/kubernetes/deployment.yml
kubectl delete -f ./comments/kubernetes/deployment-v2.yml
kubectl delete -f ./stock/kubernetes/deployment.yml
kubectl delete -f ./goods/kubernetes/deployment.yml

kubectl delete -f ./user/kubernetes/service.yml
kubectl delete -f ./comments/kubernetes/service.yml
kubectl delete -f ./stock/kubernetes/service.yml
kubectl delete -f ./goods/kubernetes/service.yml
