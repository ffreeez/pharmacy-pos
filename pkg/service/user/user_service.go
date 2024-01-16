package userservice

import (
	usermodel "pharmacy-pos/pkg/db/models/user"
	userrepo "pharmacy-pos/pkg/db/repository/user"

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
func (us *UserService) GetUserByID(id uint) (*usermodel.User, error) {
	return userrepo.GetUserByID(us.DB, id)
}

// CreateUser 创建新用户
func (us *UserService) CreateUser(user *usermodel.User) error {
	return userrepo.CreateUser(us.DB, user)
}

// ResetPassword 重置用户密码
func (us *UserService) ResetPassword(password string, id uint) error {
	return userrepo.ResetPassword(us.DB, password, id)
}

// UpdateIsAdmin 修改用户权限
func (us *UserService) UpdateIsAdmin(isadmin bool, id uint) error {
	return userrepo.UpdateIsAdmin(us.DB, isadmin, id)
}

// DeleteUserByID 根据ID删除用户
func (us *UserService) DeleteUserByID(id uint) error {
	return userrepo.DeleteUserByID(us.DB, id)
}

// AuthenticateUser 验证用户的用户名和密码
func (us *UserService) AuthenticateUser(username, password string) (*usermodel.User, error) {
	user, err := userrepo.GetUserByUserName(us.DB, username)
	if err != nil {
		return nil, err
	}

	// 检查密码是否匹配
	err = usermodel.CheckPassword(user.Password, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUserName 根据用户名获取user
func (us *UserService) GetUserByUserName(username string) (*usermodel.User, error) {
	user, err := userrepo.GetUserByUserName(us.DB, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUserInfo 获取所有的用户信息
func (us *UserService) GetAllUserInfo() ([]usermodel.User, error) {
	users, err := userrepo.GetAllUserInfo(us.DB)
	if err != nil {
		return nil, err
	}
	return users, nil
}
