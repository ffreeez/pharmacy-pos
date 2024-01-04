package usermodel

import (
	"pharmacy-pos/pkg/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	Username string
	Password string
	IsAdmin  bool
}

var c = 0

// HashPassword 将密码哈希处理
func HashPassword(password string) (string, error) {

	if c == 0 {
		config.Load()
	}

	c = config.AppConfig.Jwt.Cost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c)
	return string(bytes), err
}

// CheckPassword 检查密码是否匹配
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
