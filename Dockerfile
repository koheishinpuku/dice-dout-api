FROM golang:1.15-alpine

WORKDIR /go/src/dice-dout

COPY . .

RUN go build main.go

CMD ./main
