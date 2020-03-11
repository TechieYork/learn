#!/bin/bash

cd ./user-grpc/scripts
./start.sh

cd ../../comments-grpc/scripts
./start.sh

cd ../../stock-grpc/scripts
./start.sh

cd ../../goods-http/scripts
./start.sh

cd ../../goods-grpc/scripts
./start.sh
