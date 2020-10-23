FROM golang:1.14.10

MAINTAINER Luisa Costa<lisacosta@gmail.com>

RUN go get -u github.com/gin-gonic/gin

EXPOSE 8080

COPY app /go/

RUN go build -o app

CMD ./app


