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
	dialer *gomail.Dialer
}

func (sender *Sender) Send(email Email) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "ableimat@gmx.at")
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

func NewSender(host string, username string, password string) *Sender {
	return &Sender{dialer: gomail.NewDialer(host, 465, username, password)}
}
