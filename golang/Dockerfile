FROM golang:alpine


RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN go env

WORKDIR /go/src/app
