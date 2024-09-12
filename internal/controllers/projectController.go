package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/enums"
	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func CreateProject(c *gin.Context) {
	var req struct {
		Title        string `validate:"required"`
		Tag          string `validate:"required"`
		Description  string `validate:"required"`
		Status       string
		DeliveryDate string `validate:"required"`
	}
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	dueDate, err := time.Parse(enums.Date_format, req.DeliveryDate)
	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid date format",
		})
		return
	}

	project := models.Project{
		Title:        req.Title,
		Tag:          req.Tag,
		Description:  req.Description,
		Status:       req.Status,
		DeliveryDate: dueDate,
	}

	test := config.DB.Create(&project)
	if test.Error != nil {
		helpers.ResError(c, 400, test.Error, nil)
		return
	}
	c.JSON(201, gin.H{
		"message": "New Project Created",
		"Project": gin.H{
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
	var project models.Project

	result := config.DB.Preload("Tasks.Project").Preload("Tasks.Assigner").Preload("Tasks.AssignedTo").Preload("Tasks").First(&project, id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "No project found with this Id",
		})
		return
	}
	c.JSON(200, gin.H{
		"project": project,
	})

}

func GetAllProjects(c *gin.Context) {
	var projects []models.Project

	result := config.DB.Find(&projects)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "Error retrieving projects",
		})
		return
	}
	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func GetProjectTasks(c *gin.Context) {

}
