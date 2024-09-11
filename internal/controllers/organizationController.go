package controllers

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
	"github.com/ubaniIsaac/go-project-manager/internal/mails"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func RegisterOrganization(c *gin.Context) {
	var req struct {
		Name            string `validate:"required"`
		Email           string `validate:"required,email"`
		Size            string `validate:"required,is-valid-organization-size"`
		Industry        string `validate:"required,is-valid-industry"`
		Password        string `validate:"required"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}

	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	var existingOrganization models.Organization
	email := config.DB.Find(&existingOrganization, "email = ?", req.Email)
	if email.RowsAffected > 0 {
		c.JSON(403, gin.H{
			"message": "Account already exists with this email",
		})
		return
	}

	organization := models.Organization{
		Name:     req.Name,
		Email:    req.Email,
		Size:     req.Size,
		Industry: req.Industry,
	}

	result := config.DB.Create(&organization)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result,
		})
		return
	}

	go mails.SendWelcomeMail(
		organization.Email,
		"Welcome to GOPM",
		req.Name)

	c.JSON(201, gin.H{
		"message":      "Organization registered succefully",
		"Organization": organization,
	})
}

func InviteToOrganiztion(c *gin.Context) {
	var req struct {
		Email string `validate:"required,email"`
		Role  string `validate:"required"`
	}
	c.Bind(&req)
	organizationId := c.Param("id")

	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	email := config.DB.Find(models.User{}, "email = ?", req.Email)
	if email.RowsAffected > 0 {
		c.JSON(403, gin.H{
			"message": "Account already exists with this email",
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return
	}
	user := models.User{
		Email:    req.Email,
		Password: string(hashPassword),
		Role:     req.Role,
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result,
		})
		return
	}

	token := helpers.GenerateToken()

	b64 := string(base64.StdEncoding.EncodeToString([]byte(string(token))))

	config.DB.Create(&models.Token{UserId: user.ID, Token: token})
	var org models.Organization
	result = config.DB.First(&org, organizationId)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "No Organization found with this Id",
		})
		return
	}

	go mails.SendInviteMail(
		req.Email,
		"Invite to GOPM",
		org.Name,
		"https.gopm.com/"+b64,
	)

	c.JSON(201, gin.H{
		"message": "Invite sent to user",
		// "organization": organization,
	})
}
