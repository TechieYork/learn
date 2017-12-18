#!/bin/bash

protoc --go_out=plugin=grpc:. *.proto

