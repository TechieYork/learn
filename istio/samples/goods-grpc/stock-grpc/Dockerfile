
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/stock-grpc/bin/
COPY ./configs /app/stock-grpc/configs/

WORKDIR /app/stock-grpc/bin

ENTRYPOINT ["./stock-grpc"]
