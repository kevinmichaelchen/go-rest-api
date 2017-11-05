// main.go

package main

import "os"

func main() {
	a := App{}
	a.Initialize(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))

	a.Run(":8080")
}
