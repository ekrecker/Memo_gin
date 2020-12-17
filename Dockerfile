FROM golang:1.15-buster

RUN apt-get update && apt-get install git 

RUN mkdir /go/src/work
WORKDIR /go/src/

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/go-sql-driver/mysql
