
FROM centos:latest
MAINTAINER Techie

COPY ./build /app/user/bin/
COPY ./configs /app/user/configs/

WORKDIR /app/user/bin

ENTRYPOINT ["./user"]
