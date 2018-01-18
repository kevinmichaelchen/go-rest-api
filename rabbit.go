package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
)

func getRabbitMqPort() int {
	port, err := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		port = 5672
	}
	return port
}

var (
	rci = &RabbitConnectionInfo{
		user:         os.Getenv("RABBITMQ_USER"),
		password:     os.Getenv("RABBITMQ_PASSWORD"),
		host:         os.Getenv("RABBITMQ_HOST"),
		port:         getRabbitMqPort(),
		exchangeName: "clarakm-projects-exchange",
		routingKey:   "projects.#",
		queueName:    "clarakm-projects-queue",
	}
)

type RabbitConnectionInfo struct {
	user, password, host                string
	port                                int
	exchangeName, queueName, routingKey string
}

type RabbitListener struct {
}

type RabbitSender struct {
}

func (rci *RabbitConnectionInfo) getConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", rci.user, rci.password, rci.host, rci.port)
}

func (s *RabbitSender) send(msg string) {
	connString := rci.getConnectionString()
	log.Printf("Connecting to RabbitMQ URL: %s", connString)
	conn, err := amqp.Dial(connString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		rci.exchangeName,
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.Publish(
		rci.exchangeName,
		rci.routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", msg)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s", msg)
		panic(err)
	}
}
