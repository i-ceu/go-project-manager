package main

import (
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/routes"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	routes.RegisterRoutes()
}
