package config

import (
	"gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres@localhost:5432/postgres"))
	if err != nil {
		panic(err)

	}
	db.AutoMigrate(&models.User{})
	DB = db
}
