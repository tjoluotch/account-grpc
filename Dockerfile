# Go version
FROM golang:1.13 AS build-env

WORKDIR /grpc-acc
COPY . /grpc-acc
WORKDIR /grpc-acc/service
ENTRYPOINT go run run.go
EXPOSE 2000/tcp