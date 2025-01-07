package main

import (
	"github.com/i-ceu/go-project-manager/internal/config"
	"github.com/i-ceu/go-project-manager/internal/seeders"
)

func main() {
	// Create and run seeder
	config.LoadEnvVariables()
	config.ConnectToDB()

	seeder := seeders.NewDatabaseSeeder(config.DB)
	if err := seeder.Run(); err != nil {
		panic("Failed to seed database: " + err.Error())
	}
}
