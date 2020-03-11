
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/goods-grpc/bin/
COPY ./configs /app/goods-grpc/configs/

WORKDIR /app/goods-grpc/bin

ENTRYPOINT ["./goods-grpc"]
