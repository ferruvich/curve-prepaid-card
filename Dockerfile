FROM golang:alpine

# Preparing workdir
RUN mkdir -p /go/src/github.com/ferruvich/curve-prepaid-card
WORKDIR /go/src/github.com/ferruvich/curve-prepaid-card
ADD . /go/src/github.com/ferruvich/curve-prepaid-card

# Force the go compiler to use modules
ENV GO111MODULE=on

# Adding needed dependencies
RUN apk add --no-cache bash git gcc libc-dev

RUN go mod download