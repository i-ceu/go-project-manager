package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"github.com/i-ceu/go-project-manager/internal/services"
)

func CreateSprint(c *gin.Context) {
	var req requests.CreateSprintRequest
	c.Bind(&req)
	userID, _ := c.MustGet("userID").(string)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	projectId := c.Param("projectId")

	sprint, err := services.CreateSprint(&req, userID, projectId)
	if err != nil {
		helpers.ResError(c, 403, err.Error(), nil)
		return
	}

	helpers.Ok(c, 201, "Sprint Created", sprint)

}

// func ViewSprint(c *gin.Context){

// }
