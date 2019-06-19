FROM golang:alpine

WORKDIR /go/src/app

COPY / .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ENV GO111MODULE=on
RUN go mod download
RUN ls /go/pkg/mod -al

RUN go env


