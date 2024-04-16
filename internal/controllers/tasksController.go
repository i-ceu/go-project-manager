package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
	"gorm.io/gorm/clause"
)

func CreateTask(c *gin.Context) {
	var req struct {
		Title       string
		Description string
		DueDate     string
		Status      string
		Assigner    int64
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
		AssignerID:   userID,
		AssignedToID: int(req.AssignedTo),
	}

	newTask := config.DB.Create(&task)
	if newTask.Error != nil {
		c.JSON(400, gin.H{
			"message": newTask,
		})
		return
	}

	returnTasksDetails(&task)

	c.JSON(201, gin.H{
		"message": "Task Created",
		"Task":    task,
	})
}
func GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task

	err := checkTask(&task, id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "No Task with Id",
		})
		return
	}
	c.JSON(200, gin.H{
		"task": task,
	})
}

func AssignTask(c *gin.Context) {
	var req struct {
		AssignedTo int64
	}
	c.Bind(&req)
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task
	var user models.User

	err := checkTask(&task, id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "No Task with Id",
		})
		return
	}
	err = checkUser(&user, int(req.AssignedTo))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "No User with Id",
		})
		return
	}
	result := config.DB.Model(&task).Update("AssignedTo", &user)
	if result.Error != nil {
		c.JSON(403, gin.H{
			"message": "Error Updating Task",
			"error":   result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Task assigned to " + user.Firstname + " " + user.Lastname,
		"task":    task,
	})
}

func UpdateTask(c *gin.Context) {
	var req struct {
		Title       string
		Description string
		DueDate     string
		Status      string
	}
	c.Bind(&req)
	id, _ := strconv.Atoi(c.Param("id"))

	var task models.Task
	err := checkTask(&task, id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "No Task with Id",
		})
		return
	}
	var dueDate time.Time
	if req.DueDate != "" {
		dueDate, err = time.Parse("2006-01-02", req.DueDate)
		if err != nil {
			c.JSON(422, gin.H{
				"message": "Invalid date format",
			})
			return
		}
	}
	result := config.DB.Model(&task).Updates(models.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     dueDate,
		Status:      req.Status})
	if result.Error != nil {
		c.JSON(403, gin.H{
			"message": "Error Updating Task",
			"error":   result.Error,
		})
		return
	}

	returnTasksDetails(&task)

	c.JSON(200, gin.H{
		"message": "Task Updated",
		"task":    task,
	})

}

func checkTask(task *models.Task, id int) error {
	result := config.DB.Preload(clause.Associations).First(&task, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func checkUser(user *models.User, id int) error {
	result := config.DB.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func returnTasksDetails(t *models.Task) {
	config.DB.Model(&t).Preload("Assigner").Omit("Assigner.Password", "AssignedTo.Password").Preload("AssignedTo").Preload("Project").
		First(&t)
}
