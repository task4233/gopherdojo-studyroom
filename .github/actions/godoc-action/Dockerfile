FROM golang:latest

ENV GO111MODULE=on

COPY entrypoint.sh /entrypoint.sh
RUN chmod 700 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
