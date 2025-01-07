package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/i-ceu/go-project-manager/internal/helpers"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"github.com/i-ceu/go-project-manager/internal/services"
)

func CreateTask(c *gin.Context) {

	var req requests.CreateTaskRequest
	c.Bind(&req)
	userID, _ := c.MustGet("userID").(string)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	task, err := services.CreateTask(&req, userID)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
	}

	c.JSON(201, gin.H{
		"message": "Task Created",
		"Task":    task,
	})
}

func GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := services.GetTask(id)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}

func AssignTask(c *gin.Context) {
	var req requests.AssignTaskRequest
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	id := c.Param("id")

	task, err := services.AssignTask(&req, id)
	c.JSON(200, gin.H{
		"message": "Task assigned to " + task.AssignedTo.Firstname + " " + task.AssignedTo.Lastname,
		"task":    task,
	})
}

func UpdateTask(c *gin.Context) {
	var req requests.UpdateTaskRequest
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	id := c.Param("id")

	task, err := services.UpdateTask(&req, id)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
	}

	c.JSON(200, gin.H{
		"message": "Task Updated",
		"task":    task,
	})

}
