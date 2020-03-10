#!/bin/bash

cd ./user-grpc/scripts
./build.sh

cd ../../comments-grpc/scripts
./build.sh

cd ../../stock-grpc/scripts
./build.sh

cd ../../goods-http/scripts
./build.sh
