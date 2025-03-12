package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"go-todo-app/models"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Автоматическая миграция схемы БД
	DB.AutoMigrate(&models.Task{})
}
