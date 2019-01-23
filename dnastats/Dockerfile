FROM golang:latest

RUN go get github.com/gin-gonic/gin
RUN go get github.com/streadway/amqp
RUN go get gopkg.in/mgo.v2

# Environment variables
# ENV GIN_MODE=release
ENV MONGO_CONN_STRING=mongodb
ENV RABBIT_CONN_STRING=amqp://guest:guest@rabbitmq:5672

RUN mkdir -p /go/src/meli/dnastats
WORKDIR /go/src/meli/dnastats

COPY ./dnastats /go/src/meli/dnastats
COPY ./rabbit /go/src/meli/rabbit

RUN go build


EXPOSE 8082

CMD ["./dnastats"]
