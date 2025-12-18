package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID   string `gorm:"primaryKey" json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
