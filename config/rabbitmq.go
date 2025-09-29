package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

func InitRabbitMQ() {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		log.Fatal("RABBITMQ_URL environment variable not set")
	}

	var err error
	RabbitConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	RabbitChannel, err = RabbitConn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel:", err)
	}

	// declare exchange for notifications
	err = RabbitChannel.ExchangeDeclare(
		"notification_exchange", // name
		"topic",                 // type
		true,                    // durable
		false,                   // auto-deleted
		false,                   // internal
		false,                   // no-wait
		nil,                     // args
	)
	if err != nil {
		log.Fatal("Failed to declare exchange:", err)
	}

	log.Println("RabbitMQ connected and exchange declared")
}
