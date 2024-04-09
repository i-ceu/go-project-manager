package main

import (
	"github.com/ubaniIsaac/go-project-manager/api"
	"github.com/ubaniIsaac/go-project-manager/internal/config"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
}

func main() {
	api.RegisterRoutes()
}
