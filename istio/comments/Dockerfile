FROM centos:latest
MAINTAINER Techie York "york528@yeah.net"
ENV REFRESHED_AT 2020-03-01

COPY ./build /data/service/comments/bin/
COPY ./configs /data/service/comments/configs/

WORKDIR /data/service/comments/bin

ENTRYPOINT ["./comments"]
