package services

import (
	"errors"
	"time"

	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/enums"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
)

func CreateProject(req *requests.CreateProjectRequest) (*models.Project, error) {

	dueDate, err := time.Parse(enums.Date_format, req.DeliveryDate)
	if err != nil {
		return nil, errors.New("invalid date format")
	}

	project := models.Project{
		Title:          req.Title,
		Tag:            req.Tag,
		Description:    req.Description,
		Status:         req.Status,
		OrganizationID: req.OrganizationID,
		DeliveryDate:   dueDate,
	}

	test := config.DB.Create(&project)
	if test.Error != nil {
		return nil, test.Error
	}

	return &project, nil

}

func GetProject(id *string) (*models.Project, error) {
	var project models.Project

	result := config.DB.Preload("Tasks.Project").Preload("Tasks.Assigner").Preload("Tasks.AssignedTo").Preload("Tasks").Preload("Organization").First(&project, id)
	if result.Error != nil {
		return nil, errors.New("no project with this Id")
	}

	return &project, nil
}
func GetAllProjects(id *string) (*[]models.Project, error) {
	var projects []models.Project

	result := config.DB.Where("organization_id =?", id).Find(&projects)
	if result.Error != nil {
		return nil, errors.New("error retrirveing projects")
	}

	return &projects, nil
}
