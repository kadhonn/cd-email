package main

import (
	"github.com/kadhonn/cd-email/email"
	"github.com/kadhonn/cd-email/server"
)

func main() {
	//TODO read password from config file
	server.Run(email.NewSender(), "testpass")
}
