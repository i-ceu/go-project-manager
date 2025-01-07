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
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	project, err := services.CreateProject(&req)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
	}
	c.JSON(201, gin.H{
		"message": "New Project Created",
		"Project": gin.H{
			"Id":           project.ID,
			"Title":        project.Title,
			"Tag":          project.Tag,
			"Description":  project.Description,
			"Status":       project.Status,
			"DeliveryDate": project.DeliveryDate,
		},
	})

}

func GetProject(c *gin.Context) {
	id := c.Param("id")

	project, err := services.GetProject(&id)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
	}

	c.JSON(200, gin.H{
		"project": project,
	})

}

func GetAllProjects(c *gin.Context) { //refactor to
	id := c.Param("organizationId")

	projects, err := services.GetAllProjects(&id)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
	}
	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func GetProjectTasks(c *gin.Context) {

}
