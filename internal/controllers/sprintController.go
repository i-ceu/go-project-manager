package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"github.com/i-ceu/go-project-manager/internal/services"
)

func CreateSprint(c *gin.Context) {
	var req requests.CreateSprintRequest

	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	var user models.User
	userID := c.Param("userID")
	config.DB.Find(&user, userID).First(&user)

	sprint, err := services.CreateSprint(&req, &user)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 201, "Organization registered successfully", sprint)

}

// func ViewSprint(c *gin.Context){

// }
