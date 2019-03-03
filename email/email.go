package email

import (
	"gopkg.in/gomail.v2"
	"io"
)

type Email struct {
	Address  string
	FileName string
	Content  []byte
}

type Sender struct {
	from   string
	dialer *gomail.Dialer
}

func (sender *Sender) Send(email Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", sender.from)
	m.SetHeader("To", email.Address)
	m.SetHeader("Subject", "File from Coderdojo Linz")
	m.SetBody("text/plain", "Hello! Here is you file, have fun coding! :)")
	m.Attach(email.FileName, gomail.SetCopyFunc(copyFunc(email.Content)))

	return sender.dialer.DialAndSend(m)
}

func copyFunc(content []byte) func(writer io.Writer) error {
	return func(writer io.Writer) error {
		_, err := writer.Write(content)
		return err
	}
}

func NewSender(host string, port int, username string, password string, from string) *Sender {
	return &Sender{from: from, dialer: gomail.NewDialer(host, port, username, password)}
}
