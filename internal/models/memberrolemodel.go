package models

type MemberRole struct {
	Base
	UserID string `json:"-"`
	User   *User  `json:"user,omitempty" gorm:"constraint:OnDelete:SET NULL; foreignKey:UserID;references:ID"`
	RoleID string `json:"-"`
	Role   *Role  `json:"role,omitempty" gorm:"constraint:OnDelete:SET NULL; foreignKey:RoleID;references:ID"`
	TeamID string `json:"-"`
	Team   *Team  `json:"team,omitempty" gorm:"constraint:OnDelete:SET NULL; foreignKey:TeamID;references:ID"`
}
