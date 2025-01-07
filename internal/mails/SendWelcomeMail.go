package mails

import (
	"html/template"
	"log"

	"github.com/i-ceu/go-project-manager/internal/helpers"
)

func SendWelcomeMail(
	recipient string,
	subject string,
	name string,
	link string) error {
	values := struct {
		Name string
		Link template.URL
	}{
		Name: name,
		Link: template.URL(link),
	}
	templateFile := "../../internal/templates/welcome.html"

	err := helpers.DeliverMail(templateFile, values, recipient, subject)
	if err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)
	}

	return nil
}
