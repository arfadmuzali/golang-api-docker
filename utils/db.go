package utils

import (
	"learn/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBinit() error {
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db

	if err = DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	return nil
}
