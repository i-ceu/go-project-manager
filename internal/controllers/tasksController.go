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
	projectId := c.Param("projectId")
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	task, err := services.CreateTask(&req, userID, projectId)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
		return
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
		return
	}

	c.JSON(200, gin.H{
		"task": task,
	})
}

func GetAllTasks(c *gin.Context) {
	projectId := c.Param("projectId")

	task, err := services.GetAllTasks(projectId)
	if err != nil {
		helpers.ResError(c, 400, "error fetching project", err.Error())
		return
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
	taskId := c.Param("taskId")
	userId, _ := c.MustGet("userID").(string)

	task, err := services.AssignTask(&req, taskId, userId)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

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
	taskId := c.Param("taskId")

	task, err := services.UpdateTask(&req, taskId)
	if err != nil {
		helpers.ResError(c, 400, "error creating project", err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "Task Updated",
		"task":    task,
	})

}
