package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model     `json:"-"`
	Title          string       `json:"title" gorm:"not null"`
	Tag            string       `json:"tag" gorm:"not null"`
	Description    string       `json:"description" gorm:"not null"`
	Status         string       `json:"status" gorm:"default: inProgress"`
	DeliveryDate   time.Time    `json:"deliveryDate"`
	OrganizationID int          `json:"-"`
	Organization   Organization `gorm:"constraint:OnDelete:SET NULL;"`
	Tasks          []Task       `json:"subTasks"`
}
