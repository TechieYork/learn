FROM centos:latest
MAINTAINER Techie York "york528@yeah.net"
ENV REFRESHED_AT 2020-03-01

COPY ./build /data/service/stock/bin/
COPY ./configs /data/service/stock/configs/

WORKDIR /data/service/stock/bin

ENTRYPOINT ["./stock"]
