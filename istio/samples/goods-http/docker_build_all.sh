#!/bin/bash

cd ./user
docker build ./ --tag istio-test/user:v1

cd ../comments
docker build ./ --tag istio-test/comments:v1
docker build ./ --tag istio-test/comments:v2

cd ../stock
docker build ./ --tag istio-test/stock:v1

cd ../goods
docker build ./ --tag istio-test/goods:v1
