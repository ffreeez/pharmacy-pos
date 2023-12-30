package repository

import (
	"context"
	"pharmacy-pos/models"
	"gorm.io/gorm"
)

// UserRepository 定义了操作用户表的接口
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, userID uint) error
	Update(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, userID uint) (*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
}

// UserRepositoryImpl 实现了UserRepository接口
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository 创建一个新的UserRepository实例
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

// Create 添加一个新的用户到数据库
func (repo *UserRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	return repo.db.WithContext(ctx).Create(user).Error
}

// Delete 根据userID删除一个用户
func (repo *UserRepositoryImpl) Delete(ctx context.Context, userID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.User{}, userID).Error
}

// Update 更新用户信息
func (repo *UserRepositoryImpl) Update(ctx context.Context, user *models.User) error {
	return repo.db.WithContext(ctx).Save(user).Error
}

// FindByID 根据userID查找一个用户
func (repo *UserRepositoryImpl) FindByID(ctx context.Context, userID uint) (*models.User, error) {
	var user models.User
	err := repo.db.WithContext(ctx).First(&user, userID).Error
	return &user, err
}

// FindAll 返回所有用户的列表
func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	err := repo.db.WithContext(ctx).Find(&users).Error
	return users, err
}
