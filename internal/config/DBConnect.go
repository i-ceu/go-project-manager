package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	fmt.Println("Database connected")
	if err != nil {
		log.Fatal("failed to connect to DB")
	}

}
