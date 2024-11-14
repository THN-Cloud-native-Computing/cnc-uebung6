// MicroserviceB/main.go
package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declare a queue for this consumer
	q, err := ch.QueueDeclare(
		"",    // random queue name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	// Bind the queue to the "logs" exchange
	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key (ignored for fanout)
		"logs", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue: %s", err)
	}

	// Set up a consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	// Process messages
	go func() {
		for d := range msgs {
			log.Printf("MicroserviceB received: %s", d.Body)
			time.Sleep(1 * time.Second) // simulate work
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	select {}
}
