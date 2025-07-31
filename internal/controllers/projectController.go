package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"github.com/i-ceu/go-project-manager/internal/services"
)

func CreateProject(c *gin.Context) {
	var req requests.CreateProjectRequest
	c.Bind(&req)
	teamId := c.Param("teamId")
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	project, err := services.CreateProject(&req, teamId)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
		return
	}

	c.JSON(201, gin.H{
		"message": "New Project Created",
		"Project": gin.H{
			"id":           project.ID,
			"title":        project.Title,
			"tag":          project.Tag,
			"description":  project.Description,
			"status":       project.Status,
			"deliveryDate": project.DeliveryDate,
			"startDate":    project.StartDate,
			"workFlow":     project.WorkFlow,
		},
	})

}

func GetProject(c *gin.Context) {
	projectId := c.Param("projectId")

	project, err := services.GetProject(&projectId)
	if err != nil {
		helpers.ResError(c, 400, "false", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"project": project,
	})

}

func GetAllProjects(c *gin.Context) { //refactor to
	id := c.Param("teamId")

	projects, err := services.GetAllProjects(&id)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
		return
	}
	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func GetProjectTasks(c *gin.Context) {

}
