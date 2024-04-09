package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func CreateTask(c *gin.Context) {
	var req struct {
		Title       string
		Description string
		DueDate     string
		Status      string
		Assignee    int64
		Project     int64
		AssignedTo  int64
	}
	c.Bind(&req)

	userID, _ := strconv.Atoi(c.MustGet("userID").(string))
	dueDate, err := time.Parse("2006-01-02", req.DueDate)
	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid date format",
		})
		return
	}

	task := models.Task{
		Title:        req.Title,
		Description:  req.Description,
		Status:       req.Status,
		DueDate:      dueDate,
		ProjectID:    int(req.Project),
		AssigneeID:   userID,
		AssignedToID: int(req.AssignedTo),
	}

	newTask := config.DB.Create(&task)
	if newTask.Error != nil {
		c.JSON(400, gin.H{
			"message": newTask,
		})
		return
	}

	result := returnTasksDetails(task)

	c.JSON(201, gin.H{
		"message": "Task Created",
		"Task":    result,
	})
}
func GetTasks(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	result := config.DB.Preload("Project").Preload("Assignee").Preload("AssignedTo").First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "No task found with this Id",
		})
		return
	}
	c.JSON(200, gin.H{
		"task": task,
	})
}

func returnTasksDetails(t models.Task) struct {
	Project     string `json:"project"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"dueDate"`
	Assignee    string `json:"assignee"`
	AssignedTo  string `json:"assignedTo"`
} {
	var result struct {
		Project     string `json:"project"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		DueDate     string `json:"dueDate"`
		Assignee    string `json:"assignee"`
		AssignedTo  string `json:"assignedTo"`
	}
	config.DB.Model(&t).
		Select("tasks.title, tasks.description, tasks.status, tasks.due_date, users.firstname as assignee, projects.Title as project").
		Joins("JOIN users ON users.id = ?", t.AssigneeID).
		Joins("JOIN projects ON projects.id = ?", t.ProjectID).
		Scan(&result)
	config.DB.
		Raw("SELECT users.firstname FROM users WHERE users.id = ?", t.AssignedToID).
		Scan(&result.AssignedTo)

	return result
}
