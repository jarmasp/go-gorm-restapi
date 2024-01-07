package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(120);not null"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false"`
	UserId      uint   `json:"UserID"`
}
