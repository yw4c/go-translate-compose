FROM golang:alpine

ARG work_path

WORKDIR ${work_path}

# dockerfile path -> wokrdir
COPY / .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh tar gzip curl mysql-client

# go module
ENV GO111MODULE=on
RUN go mod download


