package main

import "github.com/asy1nc/auth-go/db"

func main() {
	db.Connect()
	db.Migrate(false)
}
