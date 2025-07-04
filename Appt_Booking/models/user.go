package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // adds ID, CreatedAt, UpdatedAt, DeletedAt fields
	ID        uint `gorm:"primaryKey"`
	Username   string `gorm:"unique" json:"username"` 
	Password   string  `json:"password"`
	IsAdmin    bool		`json:"is_Admin"` // ✅ correct format
}