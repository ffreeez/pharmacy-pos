package service

import (
	"pharmacy-pos/pkg/db/models"
	"pharmacy-pos/pkg/db/repository"

	"gorm.io/gorm"
)

// UserService 提供用户相关的服务
type UserService struct {
	DB *gorm.DB
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// GetUserByID 根据用户ID获取用户
func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return repository.GetUserByID(us.DB, id)
}

// CreateUser 创建新用户
func (us *UserService) CreateUser(user *models.User) error {
	return repository.CreateUser(us.DB, user)
}

// UpdateUser 更新用户信息
func (us *UserService) UpdateUser(user *models.User) error {
	return repository.UpdateUser(us.DB, user)
}

// DeleteUserByID 根据ID删除用户
func (us *UserService) DeleteUserByID(id uint) error {
	return repository.DeleteUserByID(us.DB, id)
}
