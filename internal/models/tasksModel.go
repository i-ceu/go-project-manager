package models

import (
	"time"
)

type Task struct {
	Base
	Title        string     `json:"title"`
	Tag          string     `json:"tag" gorm:"not null"`
	Description  string     `json:"description"`
	Status       string     `json:"status" gorm:"default: todo"`
	StartDate    *time.Time `json:"startDate" gorm:"type:datetime;default:NULL"`
	EndDate      *time.Time `json:"endDate" gorm:"type:datetime;default:NULL"`
	ProjectID    string     `json:"-"`
	Project      *Project   `json:"project,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:ProjectID;references:ID"`
	SprintID     string     `json:"-" gorm:"default:NULL"`
	Sprint       *Sprint    `json:"sprint,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:SprintID;references:ID"`
	CreatedByID  string     `json:"-"`
	CreatedBy    *User      `json:"created_by,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:CreatedByID;references:ID"`
	AssignerID   string     `json:"-" gorm:"default:NULL"`
	Assigner     *User      `json:"assigned_by" gorm:"constraint:OnDelete:SET NULL;foreignKey:AssignerID;references:ID"`
	AssignedToID string     `json:"-" gorm:"default:NULL"`
	AssignedTo   *User      `json:"assigned_to" gorm:"constraint:OnDelete:SET NULL;foreignKey:AssignedToID;references:ID"`
}
