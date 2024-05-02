package mails

import (
	"bytes"
	"log"

	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
)

func SendWelcomeMail(
	recipient string,
	subject string,
	name string) error {
	values := struct {
		Name string
	}{
		Name: name,
	}

	var tpl bytes.Buffer
	templateFile := "../../internal/templates/welcome.html"
	if err := helpers.ParseTemplate(templateFile, &tpl, values); err != nil {
		log.Fatalf("Failed to parse and execute template: %v", err)
	}
	err := helpers.DeliverMail(tpl, recipient, subject)
	if err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)

	}
	return nil
}
