FROM golang:1.15-alpine

RUN apk update && \
    apk upgrade && \
    apk add git gcc

RUN mkdir -p /go/src/work

COPY work/ /go/src/work
WORKDIR /go/src/work

RUN go get -u github.com/gin-gonic/gin \
    && go get github.com/jinzhu/gorm \
    && go get github.com/go-sql-driver/mysql 
