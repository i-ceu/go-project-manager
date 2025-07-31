package mails

import (
	"html/template"
	"log"

	"github.com/i-ceu/go-project-manager/internal/helpers"
)

func SendInviteMail(
	recipient string,
	subject string,
	teamName string,
	link string) error {
	values := struct {
		TeamName string
		Link     template.URL
	}{
		TeamName: teamName,
		Link:     template.URL(link),
	}

	templateFile := "../../internal/templates/invite.html"
	err := helpers.DeliverMail(templateFile, values, recipient, subject)
	if err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)

	}
	return nil
}
