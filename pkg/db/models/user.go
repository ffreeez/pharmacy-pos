package models

import (
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	UserName string
	Password string
	IsAdmin  bool
}
