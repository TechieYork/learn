#!/bin/bash

cd ./user-grpc
docker build ./ --tag istio-test/user-grpc:v1

cd ../comments-grpc
docker build ./ --tag istio-test/comments-grpc:v1
docker build ./ --tag istio-test/comments-grpc:v2

cd ../stock-grpc
docker build ./ --tag istio-test/stock-grpc:v1

cd ../goods-http
docker build ./ --tag istio-test/goods-http:v1

cd ../goods-grpc
docker build ./ --tag istio-test/goods-grpc:v1
