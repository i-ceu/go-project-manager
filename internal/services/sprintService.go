package services

import (
	"errors"
	"time"

	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/enums"
	"github.com/i-ceu/go-project-manager/internal/models"
	"github.com/i-ceu/go-project-manager/internal/requests"
)

func CreateSprint(req *requests.CreateSprintRequest, userId string, projectId string) (*models.Sprint, error) {
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

	var user models.User
	config.DB.Find(&user, userId).First(&user)

	sprint := models.Sprint{
		Name:        req.Name,
		StartDate:   &startDate,
		EndDate:     &endDate,
		Status:      "pending",
		CreatedByID: userId,
		ProjectID:   projectId,
	}

	spr := config.DB.Create(&sprint)
	if spr.Error != nil {
		return nil, spr.Error
	}

	return &sprint, nil
}
