package models

import "time"

type Note struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" binding:"required" gorm:"not null"`
	Content   string    `json:"content" binding:"required" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	IsPublic  *bool     `json:"is_public" binding:"required" gorm:"not null"`
}
