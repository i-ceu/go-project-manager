package mails

import (
	"fmt"
	"log"

	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
)

func SendInviteMail(
	recipient string,
	subject string,
	organizationName string,
	link string) error {
	values := struct {
		OrganizationName string
		Link             string
	}{
		OrganizationName: organizationName,
		Link:             link,
	}

	fmt.Println(link)

	templateFile := "../../internal/templates/invite.html"
	err := helpers.DeliverMail(templateFile, values, recipient, subject)
	if err != nil {
		log.Fatalf("Failed to deliver mail: %v", err)

	}
	return nil
}
