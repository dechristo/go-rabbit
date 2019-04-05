package main

import (
	"log"
	"os"
	"strings"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Unable to connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel.")
	defer ch.Close()
	
	queue, err := ch.QueueDeclare(
		"Amazing Super Fast Worker Queue", 
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Ooops! Something went wrong declaring the queue")

	user_input := os.Args
	message := strings.Join(user_input[1:], " ")

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing {
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body: []byte(message),
		},
	)
	log.Printf(" [--->] Sent %s", message)
	failOnError(err, "Publishing message failed.")	
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}