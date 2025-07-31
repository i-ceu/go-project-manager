package services

import (
	"errors"
	"time"

	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/enums"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
	"gorm.io/gorm"
)

func CreateProject(req *requests.CreateProjectRequest, teamId string) (*models.Project, error) {

	startDate, err := time.Parse(enums.Date_format, req.StartDate)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	dueDate, err := time.Parse(enums.Date_format, req.DeliveryDate)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	var team models.Team
	if err = config.DB.First(&team, "id = ?", teamId).Error; err != nil {
		return nil, errors.New("team not found")
	}

	project := models.Project{
		Title:        req.Title,
		Tag:          req.Tag,
		Description:  req.Description,
		Status:       req.Status,
		TeamID:       teamId,
		StartDate:    startDate,
		DeliveryDate: dueDate,
		WorkFlow:     models.StringWorkflow(req.WorkFlow),
	}

	test := config.DB.Create(&project)
	if test.Error != nil {
		return nil, test.Error
	}

	return &project, nil
}

func GetProject(projectId *string) (*models.Project, error) {
	var project models.Project

	result := config.DB.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Preload("AssignedTo")
	}).
		First(&project, projectId)

	if result.Error != nil {
		return nil, errors.New("no project with this Id")
	}

	return &project, nil
}

func GetAllProjects(teamId *string) (*[]models.Project, error) {
	var projects []models.Project

	result := config.DB.Where("team_id =?", teamId).Find(&projects)
	if result.Error != nil {
		return nil, errors.New("error retrieving projects")
	}

	return &projects, nil
}
