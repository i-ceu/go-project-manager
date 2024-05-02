package mails

import (
	"bytes"
	"log"

	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
)

func SendAssignTaskMail(
	recipient string,
	subject string,
	title string,
	assigner string) error {

	values := struct {
		Title    string
		Assigner string
	}{
		Title:    title,
		Assigner: assigner,
	}

	var tpl bytes.Buffer
	templateFile := "../../internal/templates/taskAssigned.html"
	if err := helpers.ParseTemplate(templateFile, &tpl, values); err != nil {
		log.Fatalf("Failed to parse and execute template: %v", err)
	}
	err := helpers.DeliverMail(tpl, recipient, subject)
	if err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)

	}
	return nil
}
