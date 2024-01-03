package repository

import (
	"pharmacy-pos/pkg/db/models"
	logger "pharmacy-pos/pkg/util"

	"gorm.io/gorm"
)

var logs = logger.GetLogger()

// GetUserByID 根据用户ID获取用户
func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
	user := &models.User{}
	result := db.First(user, id)
	if result.Error != nil {
		logs.Errorf("根据用户ID获取用户失败, ID: %d", id)
		return nil, result.Error
	}
	logs.Infof("根据用户ID获取用户成功, ID: %d", id)
	return user, nil
}

// CreateUser 创建新用户
func CreateUser(db *gorm.DB, user *models.User) error {
	result := db.Create(user)
	if result.Error != nil {
		logs.Errorf("创建新用户失败, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
		return result.Error
	}
	logs.Infof("创建新用户成功, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
	return result.Error
}

// UpdateUser 更新用户信息
func UpdateUser(db *gorm.DB, user *models.User) error {
	result := db.Save(user)
	if result.Error != nil {
		logs.Errorf("更新用户信息失败, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
		return result.Error
	}
	logs.Infof("更新用户信息成功, username: %s, password: %s, isadmin: %t", user.UserName, user.Password, user.IsAdmin)
	return result.Error
}

// DeleteUserByID 根据ID删除用户
func DeleteUserByID(db *gorm.DB, id uint) error {
	result := db.Delete(&models.User{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除用户失败, id: %d", id)
		return result.Error
	}
	logs.Infof("根据ID删除用户成功, id: %d", id)
	return result.Error
}
