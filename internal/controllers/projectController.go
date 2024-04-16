package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func CreateProject(c *gin.Context) {

	var req struct {
		Title        string
		Description  string
		Status       string
		DeliveryDate string
	}
	c.Bind(&req)

	dueDate, err := time.Parse("2006-01-02", req.DeliveryDate)
	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid date format",
		})
		return
	}

	project := models.Project{
		Title:        req.Title,
		Description:  req.Description,
		Status:       req.Status,
		DeliveryDate: dueDate,
	}
	var result struct {
		Title        string `json:"title"`
		Description  string `json:"description"`
		Status       string `json:"status"`
		DeliveryDate string `json:"deliveryDate"`
	}

	test := config.DB.Create(&project).Scan(&result)
	if test.Error != nil {
		c.JSON(400, gin.H{
			"message": test,
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "New Project Created",
		"Project": gin.H{
			"Title":        project.Title,
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
