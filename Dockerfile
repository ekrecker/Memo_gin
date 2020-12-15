FROM golang:1.15-buster

RUN apt-get update && apt-get install git 

RUN mkdir /go/src/work
WORKDIR /go/src/work

RUN go get github.com/gin-gonic/gin
