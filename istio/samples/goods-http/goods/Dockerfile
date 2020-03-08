
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/goods/bin/
COPY ./configs /app/goods/configs/

WORKDIR /app/goods/bin

ENTRYPOINT ["./goods"]
