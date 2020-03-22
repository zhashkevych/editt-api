FROM golang:1.14.0-alpine3.11 AS builder

RUN go version
RUN apk add git

COPY ./ /go/src/github.com/zhashkevych/editt-api/
WORKDIR /go/src/github.com/zhashkevych/editt-api/

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /go/src/github.com/zhashkevych/editt-api/.bin/app .
COPY --from=0 /go/src/github.com/zhashkevych/editt-api/pkg/config/ ./config/

EXPOSE 8000

CMD [ "./app"]