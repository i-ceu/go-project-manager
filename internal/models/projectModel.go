package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Status       string    `json:"status" gorm:"default: inProgress"`
	DeliveryDate time.Time `json:"deliveryDate"`
	Tasks        []Task    `json:"subTasks"`
}
