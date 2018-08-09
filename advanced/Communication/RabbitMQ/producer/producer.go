package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Starting RabbitMQ producer...")
	time.Sleep(10 * time.Second)

	// Dial connection to broker
	conn, err := amqp.Dial(brokerAddr())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Open a channel for Communication
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// delcare the queue
	q, err := ch.QueueDeclare(
		queue(), // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgCount := 0

	doneCh := make(chan struct{})

	go func() {
		for {
			msgCount++
			body := fmt.Sprintf("Hello RabbitMQ message %v", msgCount)

			// publish the message
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				},
			)
			log.Printf(" [x] Sent %s", body)
			failOnError(err, "Failed to publish the message")

			time.Sleep(5 * time.Second)
		}
	}()

	<-doneCh
}

func brokerAddr() string {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if len(brokerAddr) == 0 {
		brokerAddr = "amqp://guest:guest@localhost:5672/"
	}
	return brokerAddr
}

func queue() string {
	queue := os.Getenv("QUEUE")
	if len(queue) == 0 {
		queue = "default-queue"
	}
	return queue
}

func failOnError(err error, msg string) {
	if err != nil {

	}
}
