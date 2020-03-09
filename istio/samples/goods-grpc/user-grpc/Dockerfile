
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/user-grpc/bin/
COPY ./configs /app/user-grpc/configs/

WORKDIR /app/user-grpc/bin

ENTRYPOINT ["./user-grpc"]
