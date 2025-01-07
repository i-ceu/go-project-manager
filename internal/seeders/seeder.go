package seeders

import (
	"log"

	"gorm.io/gorm"
)

// Seeder interface for all seeders
type Seeder interface {
	Seed(*gorm.DB) error
}

// DatabaseSeeder manages all seeders
type DatabaseSeeder struct {
	db *gorm.DB
}

// NewDatabaseSeeder creates a new database seeder instance
func NewDatabaseSeeder(db *gorm.DB) *DatabaseSeeder {
	return &DatabaseSeeder{db: db}
}

// Run executes all seeders
func (s *DatabaseSeeder) Run() error {
	log.Println("Starting database seeding...")

	// Add all your seeders here
	seeders := []Seeder{
		&RoleSeeder{}, // This will be your role seeder
	}

	for _, seeder := range seeders {
		if err := seeder.Seed(s.db); err != nil {
			return err
		}
	}

	log.Println("Database seeding completed successfully")
	return nil
}

// RoleSeeder implements the Seeder interface
type RoleSeeder struct{}

// Seed implements the seeding logic
func (s *RoleSeeder) Seed(db *gorm.DB) error {
	log.Println("Running role seeder...")
	return SeedRoles(db) // This calls your existing SeedRoles function
}
