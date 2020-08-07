FROM golang:latest

ENV GO111MODULE=on

RUN apt update
RUN apt upgrade -y


ENV GOBIN /go/bin

RUN go get golang.org/x/tools/cmd/godoc

COPY entrypoint.sh /entrypoint.sh
RUN chmod 700 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
