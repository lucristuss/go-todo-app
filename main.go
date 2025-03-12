package main

import (
	"go-todo-app/database"
	"go-todo-app/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Инициализация БД
	database.InitDatabase()

	// Определяем маршруты
	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTaskByID)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	// Запуск сервера
	r.Run(":80")
}
