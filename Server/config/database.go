package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"go-todo-app/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DB.AutoMigrate(&models.Todo{})
}
