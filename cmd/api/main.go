package main

import (
	"github.com/ubaniIsaac/go-project-manager/internal/config"
	"github.com/ubaniIsaac/go-project-manager/internal/routes"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	routes.RegisterRoutes()
}
