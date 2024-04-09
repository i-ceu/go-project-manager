package main

import (
	"github.com/ubaniIsaac/go-project-manager/internal/config"
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
}
