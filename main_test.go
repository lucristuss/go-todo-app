package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setupTestDB инициализирует тестовую базу данных
func setupTestDB() *gorm.DB {
	dsn := "host=todo-db user=user password=password dbname=todo_test sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка инициализации тестовой БД: %v", err)
	}
	database.AutoMigrate(&Task{})
	return database
}

// setupRouter создает маршрутизатор с тестовой БД
func setupRouter() *gin.Engine {
	db = setupTestDB() // Подключаем тестовую БД
	router := gin.Default()

	router.GET("/tasks", GetTasks)
	router.GET("/tasks/:id", GetTaskByID)
	router.POST("/tasks", CreateTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)

	return router
}

func TestGetTasks(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	router := setupRouter()

	updatedTask := Task{Title: "Updated Task", Description: "Updated description", Completed: true}
	jsonData, _ := json.Marshal(updatedTask)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusNotFound}, w.Code)
}

func TestDeleteTask(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	router.ServeHTTP(w, req)

	assert.Contains(t, []int{http.StatusOK, http.StatusNotFound}, w.Code)
}
