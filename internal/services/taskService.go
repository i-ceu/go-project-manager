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
	"gorm.io/gorm/clause"
)

func CreateTask(req *requests.CreateTaskRequest, userID string) (*models.Task, error) {

	var project models.Project
	var sprint models.Sprint
	err := checkProject(&project, req.ProjectId)
	if err != nil {
		return nil, errors.New("no project with Id")
	}
	tag := generateTaskTag(&project, 0)

	err = checkSprint(&sprint, req.SprintID)
	if err != nil {
		return nil, errors.New("invalid sprint id")
	}

	var startDate, endDate time.Time
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

	task := models.Task{
		Title:        req.Title,
		Tag:          tag,
		Description:  req.Description,
		Status:       req.Status,
		StartDate:    startDate,
		EndDate:      endDate,
		ProjectID:    req.ProjectId,
		SprintID:     req.SprintID,
		CreatedByID:  userID,
		AssignerID:   req.Assigner,
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

func UpdateTask(req *requests.UpdateTaskRequest, id string) (*models.Task, error) {
	var task models.Task
	err := checkTask(&task, id)
	if err != nil {
		return nil, errors.New("no task with Id")
	}
	var startDate, endDate time.Time
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
	result := config.DB.Model(&task).Updates(models.Task{
		Title:       req.Title,
		Description: req.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		Status:      req.Status})
	if result.Error != nil {
		return nil, errors.New("error updating task")
	}

	getTasksDetails(&task)

	return &task, nil
}

func AssignTask(req *requests.AssignTaskRequest, id string) (*models.Task, error) {

	var task models.Task
	var user models.User

	err := checkTask(&task, id)
	if err != nil {
		return nil, err
	}

	err = checkUser(&user, req.AssignedTo)
	if err != nil {
		return nil, err
	}

	result := config.DB.Model(&task).Update("AssignedTo", &user)
	if result.Error != nil {
		return nil, result.Error
	}

	go mails.SendAssignTaskMail(
		user.Email,
		"Task Assigned to you",
		task.Title, task.Assigner.Firstname+" "+task.Assigner.Lastname)

	return &task, nil
}

func checkTask(task *models.Task, id string) error {
	result := config.DB.Preload(clause.Associations).Find(&task, id)
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
	result := config.DB.Find(&user, id)
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
	config.DB.Model(&t).Preload("Assigner").Preload("AssignedTo").Preload("Project").
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
