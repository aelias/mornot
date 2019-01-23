package rabbit

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

// DnaMessage is the struct of DNA matrix transmited by rabbitmq
type DnaMessage struct {
	IsMutant bool
	Dna      []string
}

var rabbitConnString string
var conn *amqp.Connection
var ch *amqp.Channel
var queue amqp.Queue
var err error

func init() {
	rabbitConnString = os.Getenv("RABBIT_CONN_STRING")
	if rabbitConnString == "" {
		rabbitConnString = "amqp://guest:guest@localhost:5672"
	}
	conn, err = amqp.Dial(rabbitConnString)
	// Retry if the rabbitmq server is not up and running
	retries := 0
	for err != nil && retries < 10 {
		time.Sleep(5 * time.Second)
		conn, err = amqp.Dial(rabbitConnString)
		retries++
		logOnError(err, "Failed to connect. Retrying: ")
	}
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	ch, err = conn.Channel()
	// defer ch.Close()
	queue, err = ch.QueueDeclare(
		"dnaqueue", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func logOnError(err error, msg string) {
	log.Printf("%s: %s", msg, err)
}

// Publish a message in the queue
func Publish(dna DnaMessage) {
	// body := "Hello World!"
	body, _ := json.Marshal(dna)
	err := ch.Publish(
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // inmediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})

	failOnError(err, "Failed to publish a message")
}

// Consume consumes messages from the queue
func Consume(save func(DnaMessage)) {
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	var dna DnaMessage
	for d := range msgs {
		json.Unmarshal(d.Body, &dna)
		save(dna)
		log.Printf("Received a message: %s", d.Body)
	}

}
