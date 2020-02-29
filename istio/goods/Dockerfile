FROM centos:latest
MAINTAINER Techie York "york528@yeah.net"
ENV REFRESHED_AT 2020-03-01

COPY ./build /data/service/goods/bin/
COPY ./configs /data/service/goods/configs/

WORKDIR /data/service/goods/bin

ENTRYPOINT ["./goods"]
