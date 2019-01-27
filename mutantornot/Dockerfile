FROM golang:latest

RUN go get github.com/gin-gonic/gin
RUN go get github.com/streadway/amqp
RUN	go get github.com/dghubble/sling

# Environment variables
# ENV GIN_MODE=release
ENV RABBIT_CONN_STRING=amqp://guest:guest@rabbitmq:5672

RUN mkdir -p /go/src/meli/mutantornot
WORKDIR /go/src/meli/mutantornot

COPY ./mutantornot /go/src/meli/mutantornot
COPY ./rabbit /go/src/meli/rabbit

RUN go build

EXPOSE 8081

CMD ["./mutantornot"]
