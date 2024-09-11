package controllers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/helpers"
	"github.com/ubaniIsaac/go-project-manager/internal/mails"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
	"gorm.io/gorm/clause"
)

func CreateTask(c *gin.Context) {

	var req struct {
		Title       string `validate:"required"`
		Tag         string
		Description string `validate:"required"`
		DueDate     string `validate:"required"`
		Status      string
		Assigner    string
		Project     string `validate:"required"`
		AssignedTo  string
	}
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}

	var project models.Project
	err = checkProject(&project, req.Project)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "No Project with Id",
		})
		return
	}
	tag := generateTaskTag(&project, 0)

	userID, _ := c.MustGet("userID").(string)
	dueDate, err := time.Parse("enums.Date_format", req.DueDate)
	if err != nil {
		c.JSON(422, gin.H{
			"message": "Invalid date format",
		})
		return
	}

	task := models.Task{
		Title:        req.Title,
		Tag:          tag,
		Description:  req.Description,
		Status:       req.Status,
		DueDate:      dueDate,
		ProjectID:    req.Project,
		AssignerID:   userID,
		AssignedToID: req.AssignedTo,
	}

	newTask := config.DB.Create(&task)
	if newTask.Error != nil {
		c.JSON(422, gin.H{
			"message": newTask.Error,
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
	id := c.Param("id")
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
		AssignedTo string
	}
	c.Bind(&req)
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	id := c.Param("id")
	var task models.Task
	var user models.User

	err = checkTask(&task, id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "No Task with Id",
		})
		return
	}

	err = checkUser(&user, req.AssignedTo)
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

	go mails.SendAssignTaskMail(
		user.Email,
		"Task Assigned to you",
		task.Title, task.Assigner.Firstname+" "+task.Assigner.Lastname)

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
	err := helpers.ValidateReq(req)
	if err != nil {
		c.JSON(422, gin.H{
			"message": err.Error(),
		})
		return
	}
	id := c.Param("id")

	var task models.Task
	err = checkTask(&task, id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "No Task with Id",
		})
		return
	}
	var dueDate time.Time
	if req.DueDate != "" {
		dueDate, err = time.Parse("enums.Date_format", req.DueDate)
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

func checkTask(task *models.Task, id string) error {
	result := config.DB.Preload(clause.Associations).Find(&task, id)
	if result.RowsAffected == 0 {
		return errors.New("no task with this ID")
	}
	return nil
}
func checkProject(project *models.Project, id string) error {
	result := config.DB.Preload("Tasks").Find(&project, id)
	if result.RowsAffected == 0 {
		return errors.New("No project with this ID")
	}
	return nil
}
func checkUser(user *models.User, id string) error {
	result := config.DB.Find(&user, id)
	if result.RowsAffected == 0 {
		return errors.New("now user with id")
	}
	return nil
}
func returnTasksDetails(t *models.Task) {
	config.DB.Model(&t).Preload("Assigner").Omit("Assigner.Password", "AssignedTo.Password").Preload("AssignedTo").Preload("Project").
		First(&t)
}

func generateTaskTag(project *models.Project, start int) string {
	var tasks models.Task
	projectTasks := len(project.Tasks)
	numberTag := 101 + projectTasks + start
	tag := project.Tag + "-" + strconv.Itoa(numberTag)
	result := config.DB.Where("tag = ? AND project_id = ?", tag, project.ID).Find(&tasks)
	if result.RowsAffected > 0 {
		return generateTaskTag(project, start+1)
	}
	return tag
}
