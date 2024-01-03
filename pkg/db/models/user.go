package UserModel

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户表
type User struct {
	gorm.Model
	UserName string
	Password string
	IsAdmin  bool
}

// HashPassword 将密码哈希处理
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword 检查密码是否匹配
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
