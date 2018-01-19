package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strconv"
	"time"
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

type RabbitInitializer struct {
	connection *amqp.Connection
}

type RabbitConnectionInfo struct {
	user, password, host                string
	port                                int
	exchangeName, queueName, routingKey string
}

type RabbitListener struct {
	RabbitInitializer
}

type RabbitSender struct {
	RabbitInitializer
}

func (rci *RabbitConnectionInfo) getConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", rci.user, rci.password, rci.host, rci.port)
}

func (i *RabbitInitializer) initialize() {
	// TODO the two pointers to RabbitInitializer's connection point to the same connection,
	// even though we're embedded the RabbitInitializer struct in two separate structs
	if i.connection != nil {
		log.Printf("Connection at %p is already initialized", i.connection)
		return
	}

	var conn *amqp.Connection
	var err error

	connString := rci.getConnectionString()

	log.Printf("Connecting to RabbitMQ URL: %s", connString)

	for i := 0; i < 3; i++ {
		conn, err = amqp.Dial(connString)

		if err != nil {
			log.Printf("Could not connect. Will sleep for a bit and then retry")
			time.Sleep(5 * time.Second)
		}
	}

	if conn == nil {
		failOnError(err, "Failed to connect to RabbitMQ")
	}

	i.connection = conn
}

func (l *RabbitListener) Run() {
	l.initialize()
	ch, err := l.connection.Channel()
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

	q, err := ch.QueueDeclare(
		rci.queueName, // name
		false,         // durable
		false,         // delete when unused
		true,          // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, rci.exchangeName, rci.routingKey)
	err = ch.QueueBind(
		q.Name,           // queue name
		rci.routingKey,   // routing key
		rci.exchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [RECEIVED] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func (s *RabbitSender) send(msg string) {
	s.initialize()
	ch, err := s.connection.Channel()
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
