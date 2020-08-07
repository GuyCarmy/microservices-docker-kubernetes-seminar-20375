FROM golang:1.14-alpine

ADD . /app/

WORKDIR /app

RUN go build .

ENTRYPOINT [ "/app/app" ]
