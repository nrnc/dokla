FROM debian:buster

WORKDIR /dokla

RUN apt-get update && apt-get install -y ca-certificates
RUN apt-get clean &&   rm -rf /var/lib/apt/lists/* 

ADD dokla.bin /dokla/dokla.bin

EXPOSE 9090

CMD ["/dokla/dokla.bin", "start"]