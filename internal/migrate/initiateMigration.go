package main

import (
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/migrations"
	"github.com/ubaniIsaac/go-project-manager/internal/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}
func main() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Task{})
	config.DB.AutoMigrate(&models.Project{})
	config.DB.AutoMigrate(&models.Organization{})
	config.DB.AutoMigrate(&models.Token{})
	config.DB.AutoMigrate(&models.Role{})
	config.DB.AutoMigrate(&models.StaffRole{})
	config.DB.AutoMigrate(&models.Invite{})

	migrations.DropOrganizationFromUsersTable()
}
