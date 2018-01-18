package main

import (
	"os"
	"time"
)

func main() {
	a := App{}

	// TODO this is hacky
	// I'm sleeping so we give RabbitMQ enough time to start up, lest we panic()
	// Ideally we'd be safely redialing, sleeping, and retrying a certain number of times before panicking.
	time.Sleep(10 * time.Second)

	listener := RabbitListener{}
	go listener.Run()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_DATABASE")

	a.Initialize(user, pass, host, db)

	a.Run(":8080")
}
