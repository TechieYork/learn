
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/goods-http/bin/
COPY ./configs /app/goods-http/configs/

WORKDIR /app/goods-http/bin

ENTRYPOINT ["./goods-http"]
