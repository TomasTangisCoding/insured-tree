package main

import (
	"insured/app"
	"insured/db"
)

func main() {

	db.Connect()

	app.Register()
}
