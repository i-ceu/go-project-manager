package services

import (
	"errors"
	"fmt"
	"os"

	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/mails"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
	"github.com/ubaniIsaac/go-project-manager/internal/requests"
)

func CreateOrganization(req *requests.CreateOrganizationRequest, user *models.User) (*models.Organization, error) {
	var existingOrganization models.Organization
	email := config.DB.Find(&existingOrganization, "email = ?", req.Email)
	if email.RowsAffected > 0 {
		return nil, errors.New("organization already exists with this email")
	}

	organization := models.Organization{
		Name:     req.Name,
		Email:    req.Email,
		Size:     req.Size,
		Industry: req.Industry,
	}

	org := config.DB.Create(&organization)
	if org.Error != nil {
		return nil, org.Error
	}

	staff_role := models.StaffRole{
		UserID:         user.ID,
		RoleID:         user.RoleID,
		OrganizationID: organization.ID,
	}
	sr := config.DB.Create(&staff_role)
	if sr.Error != nil {
		return nil, sr.Error
	}

	return &organization, nil
}

func InviteToOrganiztion(req *requests.SendInviteRequest, organizationId *string, user *models.User) error {
	email := config.DB.Find(models.User{}, "email = ?", req.Email)
	if email.RowsAffected > 0 {
		return errors.New("account already exists with this email")

	}

	var org models.Organization
	result := config.DB.Find(&org, organizationId)
	if result.Error != nil {
		return errors.New("No organization with this id")
	}

	invite := models.Invite{
		Firstname:      req.Firstname,
		Lastname:       req.Lastname,
		Email:          req.Email,
		RoleID:         req.RoleID,
		OrganizationID: *organizationId,
		SentByID:       user.ID,
	}
	sr := config.DB.Create(&invite)
	if sr.Error != nil {
		return sr.Error
	}

	app_url := os.Getenv("APP_URL")

	fmt.Println(app_url + invite.ID)

	go mails.SendInviteMail(
		req.Email,
		"Invite to GOPM",
		org.Name,
		app_url+invite.ID,
	)
	return nil

}
