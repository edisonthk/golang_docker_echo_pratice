FROM golang:latest

WORKDIR /opt/app

ADD . .

RUN go build