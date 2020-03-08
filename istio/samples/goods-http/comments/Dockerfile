
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/comments/bin/
COPY ./configs /app/comments/configs/

WORKDIR /app/comments/bin

ENTRYPOINT ["./comments"]
