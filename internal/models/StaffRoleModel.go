package models

type StaffRole struct {
	Base
	UserID         string       `json:"user_id"`
	User           User         `gorm:"constraint:OnDelete:SET NULL; foreignKey:UserID;references:ID"`
	RoleID         string       `json:"role_id"`
	Role           Role         `gorm:"constraint:OnDelete:SET NULL; foreignKey:RoleID;references:ID"`
	OrganizationID string       `json:"organization_id"`
	Organization   Organization `gorm:"constraint:OnDelete:SET NULL; foreignKey:OrganizationID;references:ID"`
}
