package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"github.com/i-ceu/go-project-manager/internal/services"
)

func RegisterUser(c *gin.Context) {
	var req requests.RegisterUserRequest

	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		helpers.ResError(c, 422, err.Error(), nil)
		return
	}

	user, err := services.RegisterUser(&req)

	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 201, "User registerd succefully", user)
}

func SignIn(c *gin.Context) {
	var req requests.SignInRequest

	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	user, token, err := services.SignIn(&req)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	c.JSON(200, gin.H{
		"message": "Signed in Successfuly",
		"user":    user,
		"token":   token,
	})

}

func SignInToTeam(c *gin.Context) {
	userID, _ := c.MustGet("userID").(string)
	teamId := c.Param("teamId")

	team, token, err := services.SignInToTeam(userID, teamId)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	c.JSON(200, gin.H{
		"message": "Signed in Successfuly",
		"user":    team,
		"token":   token,
	})
}

func AcceptInvite(c *gin.Context) {
	var req requests.AcceptInviteRequest
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	inviteId := c.Param("inviteId")

	user, err := services.AcceptInvite(&req, &inviteId)
	if err != nil {
		helpers.ResError(c, 404, err.Error(), nil)
		return
	}

	c.JSON(200, gin.H{
		"message": "Account created",
		"user":    *user,
	})
}

func VerifyAccount(c *gin.Context) {

	verificationId := c.Param("verificationId")
	response, _ := services.VerifyAccount(&verificationId)

	c.JSON(200, gin.H{
		"message": response,
	})
}
