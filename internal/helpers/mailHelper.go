package helpers

import (
	"bytes"
	"html/template"
	"os"

	"github.com/go-mail/mail"
)

func DeliverMail(tpl bytes.Buffer, recipient string, subject string) error {
	m := mail.NewMessage()

	m.SetHeader("From", os.Getenv("app_mail"))

	m.SetHeader("To", recipient)

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", tpl.String())

	d := mail.NewDialer(os.Getenv("smtpHost"), 2525, os.Getenv("username"), os.Getenv("password"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func ParseTemplate(templateFile string, tpl *bytes.Buffer, values interface{}) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(tpl, values); err != nil {
		return err
	}

	return nil
}
