package repository

import (
	"errors"
	UserModel "pharmacy-pos/pkg/db/models"
	logger "pharmacy-pos/pkg/util/logger"

	"gorm.io/gorm"
)

var logs = logger.GetLogger()

// GetUserByID 根据用户ID获取用户
func GetUserByID(db *gorm.DB, id uint) (*UserModel.User, error) {
	user := &UserModel.User{}
	result := db.First(user, id)
	if result.Error != nil {
		logs.Errorf("根据用户ID获取用户失败, ID: %d", id)
		return nil, result.Error
	}
	logs.Infof("根据用户ID获取用户成功, ID: %d", id)
	return user, nil
}

// CreateUser 创建新用户
func CreateUser(db *gorm.DB, user *UserModel.User) error {

	var existingUser UserModel.User
	result := db.Where("user_name = ?", user.UserName).First(&existingUser)
	if result.Error == nil {
		logs.Errorf("用户名已存在, username: %s", user.UserName)
		result.Error = errors.New("用户名已存在")
		return result.Error
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logs.Errorf("发生其他错误")
		return result.Error
	}

	var HashPassworderr error
	user.Password, HashPassworderr = UserModel.HashPassword(user.Password)
	if HashPassworderr != nil {
		logs.Errorf("处理用户密码失败")
		return HashPassworderr
	}
	result = db.Create(user)
	if result.Error != nil {
		logs.Errorf("创建新用户失败, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
		return result.Error
	}
	logs.Infof("创建新用户成功, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
	return result.Error
}

// UpdateUser 更新用户信息
func UpdateUser(db *gorm.DB, user *UserModel.User) error {
	var hasherr error
	user.Password, hasherr = UserModel.HashPassword(user.Password)
	if hasherr != nil {
		logs.Errorf("更新用户信息时，生成hash密码失败, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
		return hasherr
	}

	result := db.Model(user).Omit("CreatedAt").Save(user)
	if result.Error != nil {
		logs.Errorf("更新用户信息失败, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
		return result.Error
	}
	logs.Infof("更新用户信息成功, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
	return result.Error
}

// DeleteUserByID 根据ID删除用户
func DeleteUserByID(db *gorm.DB, id uint) error {
	result := db.Delete(&UserModel.User{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除用户失败, id: %d", id)
		return result.Error
	}
	logs.Infof("根据ID删除用户成功, id: %d", id)
	return result.Error
}

// GetUserByUserName 根据用户名查找用户
func GetUserByUserName(db *gorm.DB, username string) (*UserModel.User, error) {
	user := &UserModel.User{}
	result := db.Where("user_name = ?", username).First(user)
	if result.Error != nil {
		logs.Errorf("根据用户名查找用户失败, username: %s", username)
		return nil, result.Error
	}
	logs.Infof("根据用户名查找用户成功, username: %s", username)
	return user, nil
}
