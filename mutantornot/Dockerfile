FROM golang:latest

RUN go get github.com/gin-gonic/gin
RUN go get github.com/streadway/amqp

# ENV GIN_MODE=release
RUN mkdir -p /go/src/meli/mutantornot
WORKDIR /go/src/meli/mutantornot

COPY ./mutantornot /go/src/meli/mutantornot
COPY ./rabbit /go/src/meli/rabbit

RUN go build

EXPOSE 8081

CMD ["./mutantornot"]
