package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"github.com/i-ceu/go-project-manager/internal/services"
)

func RegisterTeam(c *gin.Context) {
	var req requests.CreateTeamRequest

	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		helpers.ResError(c, 422, err.Error(), nil)
		return
	}
	var user models.User
	userID, _ := c.MustGet("userID").(string)
	config.DB.Where("id=?", userID).First(&user)

	team, err := services.CreateTeam(&req, &user)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 201, "Team registered successfully", team)
}

func InviteToTeam(c *gin.Context) {
	var req requests.SendInviteRequest
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		helpers.ResError(c, 422, err.Error(), nil)
		return
	}
	teamId := c.Param("teamId")
	userID, exists := c.Get("userID")

	if !exists {
		helpers.ResError(c, 403, "unauthorized", nil)
		return
	}

	var user models.User
	config.DB.Where("id=?", userID).First(&user)

	fail := services.InviteToTeam(&req, &teamId, &user)
	if fail != nil {
		helpers.ResError(c, 403, fail.Error(), nil)
		return
	}

	helpers.Ok(c, 201, "Invite sent to user", nil)
}

func GetAllUserTeams(c *gin.Context) {
	userID, _ := c.MustGet("userID").(string)

	teams, err := services.GetUserTeams(userID)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 200, "All Users teams", teams)
}

func GetRoles(c *gin.Context) {
	teams, err := services.GetRoles()
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 200, "all roles returned", teams)
}

func GetTeamMembers(c *gin.Context) {
	// userID, _ := c.MustGet("userID").(string)
	teamId := c.Param("teamId")

	teams, err := services.GetTeamMembers(teamId)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 200, "All teams members", teams)
}
