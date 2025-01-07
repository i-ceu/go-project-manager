package models

import (
	"time"
)

type Task struct {
	Base
	Title        string    `json:"title"`
	Tag          string    `json:"tag" gorm:"not null"`
	Description  string    `json:"description"`
	Status       string    `json:"status" gorm:"default: todo"`
	StartDate    time.Time `json:"startDate" gorm:"type:datetime;default:NULL"`
	EndDate      time.Time `json:"endDate" gorm:"type:datetime;default:NULL"`
	ProjectID    string    `json:"-"`
	Project      Project   `gorm:"constraint:OnDelete:SET NULL;"`
	SprintID     string    `json:"-" gorm:"default:NULL"`
	Sprint       Project   `gorm:"constraint:OnDelete:SET NULL;"`
	CreatedByID  string    `json:"-"`
	CreatedBy    User      `gorm:"constraint:OnDelete:SET NULL;"`
	AssignerID   string    `json:"-" gorm:"default:NULL"`
	Assigner     User      `gorm:"constraint:OnDelete:SET NULL;"`
	AssignedToID string    `json:"-" gorm:"default:NULL"`
	AssignedTo   User      `gorm:"constraint:OnDelete:SET NULL;"`
}
