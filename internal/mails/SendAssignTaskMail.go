package mails

import (
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

	templateFile := "../../internal/templates/taskAssigned.html"
	err := helpers.DeliverMail(templateFile, values, recipient, subject)
	if err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)

	}
	return nil
}
