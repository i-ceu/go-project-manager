package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// StringWorkflow is a custom type for storing workflow as JSON array of strings
type StringWorkflow []string

// Scan implements the Scanner interface for database reads
func (w *StringWorkflow) Scan(value interface{}) error {
	if value == nil {
		*w = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot scan non-[]byte value into StringWorkflow")
	}

	return json.Unmarshal(bytes, w)
}

// Value implements the Valuer interface for database writes
func (w StringWorkflow) Value() (driver.Value, error) {
	if w == nil {
		return nil, nil
	}
	return json.Marshal(w)
}

type Project struct {
	Base
	Title            string         `json:"title" gorm:"not null"`
	Tag              string         `json:"tag" gorm:"not null"`
	Description      string         `json:"description" gorm:"not null"`
	Status           string         `json:"status" gorm:"default: in-progress"`
	StartDate        *time.Time     `json:"startDate"`
	DeliveryDate     *time.Time     `json:"deliveryDate"`
	WorkFlow         StringWorkflow `json:"workflow" gorm:"type:text"`
	TeamID           string         `json:"-"`
	Team             *Team          `json:"team,omitempty" gorm:"constraint:OnDelete:SET NULL;"`
	Tasks            []Task         `json:"tasks"`
	TasksAIGenerated bool           `json:"tasks_ai_generated" gorm:"default: 0"`
}
