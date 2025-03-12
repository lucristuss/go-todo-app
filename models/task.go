package models

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    Title       string `json:"title" binding:"required"`
    Description string `json:"description"`
    Completed   bool   `json:"completed"`
}
