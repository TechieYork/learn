FROM centos:latest
MAINTAINER Techie York "york528@yeah.net"
ENV REFRESHED_AT 2020-03-01

COPY ./build /data/service/user/bin/
COPY ./configs /data/service/user/configs/

WORKDIR /data/service/user/bin

ENTRYPOINT ["./user"]
