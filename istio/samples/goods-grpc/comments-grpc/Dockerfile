
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/comments-grpc/bin/
COPY ./configs /app/comments-grpc/configs/

WORKDIR /app/comments-grpc/bin

ENTRYPOINT ["./comments-grpc"]
