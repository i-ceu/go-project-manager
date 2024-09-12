package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
	"github.com/ubaniIsaac/go-project-manager/internal/requests"
	"github.com/ubaniIsaac/go-project-manager/internal/services"
)

func RegisterOrganization(c *gin.Context) {
	var req requests.CreateOrganizationRequest

	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		helpers.ResError(c, 422, err.Error(), nil)
		return
	}
	var user models.User
	userID := c.Param("userID")
	config.DB.Find(&user, userID).First(&user)

	organization, err := services.CreateOrganization(&req, &user)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 201, "Organization registered successfully", organization)
}

func InviteToOrganiztion(c *gin.Context) {
	var req requests.SendInviteRequest
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		helpers.ResError(c, 422, err.Error(), nil)
		return
	}
	organizationId := c.Param("id")
	userID := c.Param("userID")
	var user models.User
	config.DB.Find(&user, userID).First(&user)

	fail := services.InviteToOrganiztion(&req, &organizationId, &user)
	if fail != nil {
		helpers.ResError(c, 403, fail.Error(), nil)
		return
	}

	helpers.Ok(c, 201, "Invite sent to user", nil)
}
