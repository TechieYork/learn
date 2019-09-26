#!/bin/bash

protoc --proto_path=${GOPATH}/src --go_out=${GOPATH}/src ${GOPATH}/src/learn/test_protobuf/proto_import/protocol/proto_main/student.proto
