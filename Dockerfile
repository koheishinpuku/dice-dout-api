FROM golang:1.15-alpine

WORKDIR /go/src//github.com/koheishinpuku/dice-dout-api

COPY . .

RUN go build main.go

CMD ./app
