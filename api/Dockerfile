FROM golang:1.11

ADD . /go/src/go-poll/api

RUN go get github.com/gorilla/mux
RUN go get github.com/go-redis/redis

ENTRYPOINT go run main.go