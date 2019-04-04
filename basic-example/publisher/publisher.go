package main

import (	
	"log"

	"github.com/streadway/amqp"
)

func main () {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Unable to connect to RabbitMQ.")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Unable to open a channel.")
	defer ch.Close()
	
	queue, err := ch.QueueDeclare(
		"Amazing Super Fast Queue", 
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Ooops! Something went wrong declaring the queue")

	body := "Try not. Do or do not. There is no try."
	
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing {
			ContentType: "application/json",
			Body: []byte(body),
		},
	)
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Publishing message failed.")	
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}