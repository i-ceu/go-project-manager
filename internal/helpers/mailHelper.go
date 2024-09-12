package helpers

import (
	"bytes"
	"html/template"
	"log"
	"os"
	"strconv"

	"github.com/go-mail/mail"
)

func DeliverMail(
	templateFile string,
	values interface{},
	recipient string,
	subject string) error {

	var tpl bytes.Buffer
	if err := parseTemplate(templateFile, &tpl, values); err != nil {
		log.Fatalf("Failed to parse and execute template: %v", err)
		return err
	}

	m := mail.NewMessage()

	m.SetHeader("From", os.Getenv("app_mail"))

	m.SetHeader("To", recipient)

	m.SetHeader("Subject", subject)

	m.SetBody("text/html", tpl.String())

	port, _ := strconv.Atoi(os.Getenv("smtpport"))

	d := mail.NewDialer(os.Getenv("smtpHost"), port, os.Getenv("username"), os.Getenv("password"))

	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)
		return err
	}
	return nil
}

func parseTemplate(templateFile string, tpl *bytes.Buffer, values interface{}) error {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(tpl, values); err != nil {
		log.Fatalf("Failed to parse and execute template: %v", err)
	}

	return nil
}
