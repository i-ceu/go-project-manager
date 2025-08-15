package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/enums"
	"github.com/i-ceu/go-project-manager/internal/mails"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"gorm.io/gorm"
)

func CreateTask(req *requests.CreateTaskRequest, userID string, projectId string) (*models.Task, error) {

	var project models.Project
	var sprint models.Sprint
	err := checkProject(&project, projectId)
	if err != nil {
		return nil, errors.New("no project with Id")
	}
	tag := generateTaskTag(&project, 0)

	if len(req.SprintID) != 0 {
		err = checkSprint(&sprint, req.SprintID)
	}
	if err != nil {
		return nil, errors.New("invalid sprint id")
	}

	var startDate, endDate *time.Time
	if len(req.StartDate) != 0 {
		parsedStart, err := time.Parse(enums.Date_format, req.StartDate)
		if err != nil {
			return nil, errors.New("invalid date format")
		}
		startDate = &parsedStart
	}

	if len(req.EndDate) != 0 {
		parsedEnd, err := time.Parse(enums.Date_format, req.EndDate)
		if err != nil {
			return nil, errors.New("invalid date format")
		}
		startDate = &parsedEnd
	}

	var status string
	if len(req.Status) == 0 {
		status = project.WorkFlow[0]
	} else {
		status = req.Status
	}

	task := models.Task{
		Title:        req.Title,
		Tag:          tag,
		Description:  req.Description,
		Status:       status,
		ProjectID:    projectId,
		StartDate:    startDate,
		EndDate:      endDate,
		SprintID:     req.SprintID,
		CreatedByID:  userID,
		AssignedToID: req.AssignedTo,
	}

	newTask := config.DB.Create(&task)
	if newTask.Error != nil {
		return nil, errors.New(newTask.Error.Error())
	}

	getTasksDetails(&task)

	return &task, nil
}

func GetTask(id string) (*models.Task, error) {
	var task models.Task

	err := checkTask(&task, id)
	if err != nil {
		return nil, errors.New("no task with this id")
	}

	return &task, nil
}

func GetAllTasks(projectId string, userId string) (*[]models.Task, error) {
	var tasks []models.Task
	var result *gorm.DB
	if userId == "unassigned" {
		result = config.DB.Preload("Sprint").Preload("AssignedTo").Where("project_id = ?", projectId).
			Where("assigned_to_id IS NULL").
			Find(&tasks)
	} else if userId != "" {
		result = config.DB.Preload("Sprint").Preload("AssignedTo").Where("project_id = ? AND assigned_to_id = ?", projectId, userId).Find(&tasks)
	} else {
		result = config.DB.Preload("Sprint").Preload("AssignedTo").Where("project_id = ?", projectId).Find(&tasks)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &tasks, nil
}

func UpdateTask(req *requests.UpdateTaskRequest, taskId string) (*models.Task, error) {
	var existingTask models.Task
	if err := config.DB.First(&existingTask, "id = ?", taskId).Error; err != nil {
		return nil, errors.New("no task with Id")
	}

	var startDate, endDate time.Time
	var err error

	if len(req.StartDate) != 0 {
		startDate, err = time.Parse(enums.Date_format, req.StartDate)
		if err != nil {
			return nil, errors.New("invalid date format")
		}
	}

	if len(req.EndDate) != 0 {
		endDate, err = time.Parse(enums.Date_format, req.EndDate)
		if err != nil {
			return nil, errors.New("invalid date format")
		}
	}

	// Create a map for updates to avoid zero-value issues
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if !startDate.IsZero() {
		updates["start_date"] = startDate
	}
	if !endDate.IsZero() {
		updates["end_date"] = endDate
	}

	// Update using map instead of struct
	if err := config.DB.Model(&existingTask).Updates(updates).Error; err != nil {
		return nil, errors.New("error updating task")
	}

	// Get fresh instance with preloaded associations
	var updatedTask models.Task
	if err := config.DB.Preload("Assigner").Preload("AssignedTo").Preload("Project").
		First(&updatedTask, "id = ?", taskId).Error; err != nil {
		return nil, errors.New("error loading updated task")
	}

	return &updatedTask, nil
}

func AssignTask(req *requests.AssignTaskRequest, taskId string, userId string) (*models.Task, error) {

	var task models.Task
	var user models.User
	var assigner models.User

	err := checkTask(&task, taskId)
	if err != nil {
		return nil, err
	}

	err = checkUser(&user, req.AssignedTo)
	if err != nil {
		return nil, err
	}

	err = checkUser(&assigner, userId)
	if err != nil {
		return nil, err
	}

	result := config.DB.Model(&models.Task{}).Where("id = ?", taskId).Updates(map[string]interface{}{
		"assigned_to_id": user.ID,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	err = config.DB.Preload("Assigner").Preload("AssignedTo").Where("id = ?", taskId).First(&task).Error
	if err != nil {
		return nil, err
	}

	assignerName := assigner.Firstname + " " + assigner.Lastname

	go mails.SendAssignTaskMail(
		user.Email,
		"Task Assigned to you",
		task.Title, assignerName)

	return &task, nil
}

func checkTask(task *models.Task, id string) error {
	result := config.DB.Preload("Assigner").Preload("AssignedTo").Preload("Project").Preload("CreatedBy").Preload("Sprint").Where("id = ?", id).First(&task)
	if result.RowsAffected == 0 {
		return errors.New("no task with this ID")
	}
	return nil
}
func checkProject(project *models.Project, id string) error {
	result := config.DB.Preload("Tasks").Where("id=?", id).First(&project)
	if result.RowsAffected == 0 {
		return errors.New("no project with this ID")
	}
	return nil
}
func checkUser(user *models.User, id string) error {
	result := config.DB.Where("id = ?", id).First(&user)
	if result.RowsAffected == 0 {
		return errors.New("now user with id")
	}
	return nil
}
func checkSprint(sprint *models.Sprint, id string) error {
	result := config.DB.Find(&sprint, id)
	if result.RowsAffected == 0 {
		return errors.New("now user with id")
	}
	return nil
}
func getTasksDetails(t *models.Task) {
	config.DB.Preload("Assigner").Preload("AssignedTo").Preload("Project").
		First(t, "id = ?", t.ID)
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
