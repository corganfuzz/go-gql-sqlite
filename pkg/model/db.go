package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("portfolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate all models
	DB.AutoMigrate(&Project{}, &Skill{}, &Experience{})
}
