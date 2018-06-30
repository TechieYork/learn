#!/bin/bash

./test_docker_leadership --node="node-1" --kv="zk" --address="127.0.0.1:2181" &
./test_docker_leadership --node="node-2" --kv="zk" --address="127.0.0.1:2181" &
./test_docker_leadership --node="node-3" --kv="zk" --address="127.0.0.1:2181" &
