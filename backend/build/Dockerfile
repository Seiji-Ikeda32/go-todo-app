FROM golang:1.17-alpine

RUN apk update && apk add git

WORKDIR  /go/src/github.com/Seiji-Ikeda32/go-todo-app/backend
ADD . /go/src/github.com/Seiji-Ikeda32/go-todo-app/backend

RUN go mod download


