
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/stock/bin/
COPY ./configs /app/stock/configs/

WORKDIR /app/stock/bin

ENTRYPOINT ["./stock"]
