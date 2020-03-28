FROM golang:1.14.0-alpine3.11 AS builder

RUN go version
RUN apk add git
RUN apk --no-cache add ca-certificates

RUN go get -u github.com/go-delve/delve/cmd/dlv

WORKDIR /root/

COPY ./.bin/app .
COPY pkg/config ./config/