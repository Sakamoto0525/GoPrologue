FROM golang:1.15-alpine

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

RUN apk add --no-cache alpine-sdk git && go get -u github.com/cosmtrek/air

EXPOSE 8080
