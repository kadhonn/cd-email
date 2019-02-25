package main

import (
	"flag"
	"github.com/kadhonn/cd-email/email"
	"github.com/kadhonn/cd-email/server"
	"github.com/vharitonsky/iniflags"
	"log"
)

var (
	password     = flag.String("password", "", "the password needed in the ui")
	smtpHost     = flag.String("smtpHost", "", "the smtpHost needed to send emails")
	smtpUser     = flag.String("smtpUser", "", "the smtpUser needed to send emails")
	smtpPassword = flag.String("smtpPassword", "", "the smtpPassword needed to send emails")
)

func main() {
	iniflags.SetConfigFile("/etc/cd-email/config.ini")
	iniflags.SetAllowMissingConfigFile(true)
	iniflags.Parse()

	checkValues()

	server.Run(email.NewSender(*smtpHost, *smtpUser, *smtpPassword), *password)
}

func checkValues() {
	if *password == "" {
		log.Fatal("needs a password set")
	}
	if *smtpHost == "" {
		log.Fatal("needs a smtpHost set")
	}
	if *smtpUser == "" {
		log.Fatal("needs a smtpUser set")
	}
	if *smtpPassword == "" {
		log.Fatal("needs a smtpPassword set")
	}
}
