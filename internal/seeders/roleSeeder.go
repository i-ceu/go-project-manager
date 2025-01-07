package seeders

import (
	"log"

	"github.com/i-ceu/go-project-manager/internal/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	// Define the roles to be seeded
	roles := []models.Role{
		{
			Name: "super-admin",
		},
		{
			Name: "admin",
		},
		{
			Name: "employee",
		},
	}

	for _, role := range roles {
		var existing models.Role
		result := db.Where("name = ?", role.Name).First(&existing)

		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("Creating role: %s\n", role.Name)
			if err := db.Create(&role).Error; err != nil {
				return err
			}
		} else if result.Error != nil {
			return result.Error
		} else {
			log.Printf("Role already exists: %s\n", role.Name)
		}
	}

	return nil
}
