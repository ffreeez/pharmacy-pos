package userservice

import (
	UserModel "pharmacy-pos/pkg/db/models"
	UserRepo "pharmacy-pos/pkg/db/repository"

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
func (us *UserService) GetUserByID(id uint) (*UserModel.User, error) {
	return UserRepo.GetUserByID(us.DB, id)
}

// CreateUser 创建新用户
func (us *UserService) CreateUser(user *UserModel.User) error {
	return UserRepo.CreateUser(us.DB, user)
}

// UpdateUser 更新用户信息
func (us *UserService) UpdateUser(user *UserModel.User) error {
	return UserRepo.UpdateUser(us.DB, user)
}

// DeleteUserByID 根据ID删除用户
func (us *UserService) DeleteUserByID(id uint) error {
	return UserRepo.DeleteUserByID(us.DB, id)
}

// AuthenticateUser 验证用户的用户名和密码
func (us *UserService) AuthenticateUser(username, password string) (*UserModel.User, error) {
	user, err := UserRepo.GetUserByUserName(us.DB, username)
	if err != nil {
		return nil, err
	}

	// 检查密码是否匹配
	err = UserModel.CheckPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUserName 根据用户名获取user
func (us *UserService) GetUserByUserName(username string) (*UserModel.User, error) {
	user, err := UserRepo.GetUserByUserName(us.DB, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
