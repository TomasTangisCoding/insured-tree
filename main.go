package main

import (
	"insured/app"
	"insured/db"
)

func main() {
	// Connect to the database
	db.Connect()

	// Register the routes
	app.Register()
}
