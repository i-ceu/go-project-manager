package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal("failed to connect to DB")
	}

}
