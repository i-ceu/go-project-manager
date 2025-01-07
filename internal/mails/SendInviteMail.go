package mails

import (
	"fmt"
	"html/template"
	"log"

	"github.com/i-ceu/go-project-manager/internal/helpers"
)

func SendInviteMail(
	recipient string,
	subject string,
	organizationName string,
	link string) error {
	values := struct {
		OrganizationName string
		Link             template.URL
	}{
		OrganizationName: organizationName,
		Link:             template.URL(link),
	}

	fmt.Println(link)

	templateFile := "../../internal/templates/invite.html"
	err := helpers.DeliverMail(templateFile, values, recipient, subject)
	if err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)

	}
	return nil
}
