package main

import (
	"cd-email/email"
	"cd-email/server"
)

func main() {
	//TODO read password from config file
	server.Run(email.NewSender(), "testpass")
}
