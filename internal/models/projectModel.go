package models

import (
	"time"
)

type Project struct {
	Base
	Title          string       `json:"title" gorm:"not null"`
	Tag            string       `json:"tag" gorm:"not null"`
	Description    string       `json:"description" gorm:"not null"`
	Status         string       `json:"status" gorm:"default: inProgress"`
	DeliveryDate   time.Time    `json:"deliveryDate"`
	OrganizationID string       `json:"-"`
	Organization   Organization `gorm:"constraint:OnDelete:SET NULL;"`
	Tasks          []Task       `json:"subTasks"`
}
