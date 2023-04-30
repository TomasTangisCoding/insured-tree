package main

import (
	"insured/app"
	"insured/initiate"
)

func main() {

	initiate.SetConfig()
	initiate.ConnectDB()

	app.Register()
}
