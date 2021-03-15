FROM golang:1.15-alpine

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work

EXPOSE 8080

CMD ["go", "run", "main.go"]
