package main

import (
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/migrations"
	"github.com/i-ceu/go-project-manager/internal/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}
func main() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Task{})
	config.DB.AutoMigrate(&models.Project{})
	config.DB.AutoMigrate(&models.Team{})
	config.DB.AutoMigrate(&models.Token{})
	config.DB.AutoMigrate(&models.Role{})
	config.DB.AutoMigrate(&models.MemberRole{})
	config.DB.AutoMigrate(&models.Invite{})
	config.DB.AutoMigrate(&models.Sprint{})

	migrations.DropTeamFromUsersTable()
	migrations.DropEmailFromTeamTable()
}
