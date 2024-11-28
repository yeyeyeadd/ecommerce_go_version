package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DB_DSN")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automatically migrate models
	err = DB.AutoMigrate(&User{}, &Product{}, &Order{}, &OrderItem{}, &Review{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
