FROM golang:alpine

ARG work_path

WORKDIR ${work_path}

# dockerfile path -> wokrdir
COPY / .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh tar gzip curl mysql-client

# go module
#ENV GO111MODULE=on
#RUN go mod download

# golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.4.0/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate.linux-amd64 /bin/migrate



