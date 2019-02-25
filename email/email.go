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

//TODO properties
func NewSender() *Sender {
	return &Sender{dialer: gomail.NewDialer("mail.gmx.net", 465, "ableimat@gmx.at", "Mt6w1zUYdUvBGje9pflm")}
}
