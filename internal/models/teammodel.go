package models

type Team struct {
	Base
	Name        string `json:"teamName" gorm:"not null"`
	Description string `json:"description" `
	Size        string `json:"size"`
	Industry    string `json:"industry"`
	CreatorID   string `json:"creator_id"`
	Creator     *User  `json:"creator,omitempty" gorm:"constraint:OnDelete:SET NULL;foreignKey:CreatorID;references:ID"`
}
