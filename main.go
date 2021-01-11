package main

import (
	"web_base/app"
)

func main() {
	server := app.NewServer()
	server.RunServer()
}
