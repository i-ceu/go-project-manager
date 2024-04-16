package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Status       string    `json:"status" gorm:"default: todo"`
	DueDate      time.Time `json:"dueDate"`
	ProjectID    int       `json:"projectId"`
	Project      Project   `gorm:"constraint:OnDelete:SET NULL;"`
	AssignerID   int       `json:"assignerId"`
	Assigner     User      `gorm:"constraint:OnDelete:SET NULL;"`
	AssignedToID int       `json:"assignedToId" gorm:"default: 1"`
	AssignedTo   User      `gorm:"constraint:OnDelete:SET NULL;"`
}
