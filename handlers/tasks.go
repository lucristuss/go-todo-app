package handlers

import (
	"net/http"
	"go-todo-app/database"
	"go-todo-app/models"

	"github.com/gin-gonic/gin"
)

// Получение всех задач
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

// Получение задачи по ID
func GetTaskByID(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Создание новой задачи
func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

// Обновление задачи
func UpdateTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// Удаление задачи
func DeleteTask(c *gin.Context) {
	var task models.Task
	id := c.Param("id")

	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}
