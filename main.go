package main

import (
	"os"
)

func main() {
	a := App{}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_DATABASE")

	a.Initialize(user, pass, host, db)

	a.Run(":8080")
}
