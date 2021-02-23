FROM golang:1.15-alpine

WORKDIR /go/src/github.com/koheishinpuku/dice-dout-api

RUN go get -u bitbucket.org/liamstask/goose/cmd/goose

COPY . .

CMD go build main.go && ./main
