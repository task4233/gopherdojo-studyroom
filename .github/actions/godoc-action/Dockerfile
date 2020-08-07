FROM golang:latest

ENV GO111MODULE=on

RUN apt update
RUN apt upgrade -y

COPY entrypoint.sh /entrypoint.sh
RUN chmod 700 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
