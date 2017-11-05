// main.go

package main

import (
	"os"
	"fmt"
)

func main() {
	a := App{}

	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	db := os.Getenv("MYSQL_DATABASE")

	a.Initialize(user, pass, host, db)

	fmt.Sprintf("%s:%s %s/%s", user, pass, host, db)
	a.Run(":8080")
}
