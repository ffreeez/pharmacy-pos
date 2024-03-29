package userrepo

import (
	"errors"
	usermodel "pharmacy-pos/pkg/db/models/user"
	logger "pharmacy-pos/pkg/util/logger"

	"pharmacy-pos/pkg/util/e"

	"gorm.io/gorm"
)

var logs = logger.GetLogger()

// GetUserByID 根据用户ID获取用户
func GetUserByID(db *gorm.DB, id uint) (*usermodel.User, error) {
	user := &usermodel.User{}
	result := db.First(user, id)
	if result.Error != nil {
		logs.Errorf("根据用户ID获取用户失败, ID: %d", id)
		return nil, result.Error
	}
	logs.Infof("根据用户ID获取用户成功, ID: %d", id)
	return user, nil
}

// CreateUser 创建新用户
func CreateUser(db *gorm.DB, user *usermodel.User) error {

	var existingUser usermodel.User
	result := db.Where("user_name = ?", user.UserName).First(&existingUser)
	if result.Error == nil {
		logs.Errorf("用户名已存在, username: %s", user.UserName)
		result.Error = e.ErrUsernameExists
		return result.Error
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		logs.Errorf("发生其他错误")
		return result.Error
	}

	var HashPassworderr error
	user.Password, HashPassworderr = usermodel.HashPassword(user.Password)
	if HashPassworderr != nil {
		logs.Errorf("处理用户密码失败")
		return HashPassworderr
	}
	result = db.Create(user)
	if result.Error != nil {
		logs.Errorf("创建新用户失败, username: %s, isadmin: %t", user.UserName, user.IsAdmin)
		return result.Error
	}
	logs.Infof("创建新用户成功, username: %s, isadmin: %t", user.UserName, user.IsAdmin)
	return result.Error
}

// ResetPassword 重设用户密码
func ResetPassword(db *gorm.DB, password string, id uint) error {
	var user usermodel.User
	if err := db.First(&user, id).Error; err != nil {
		logs.Errorf("未找到用户: %d, error: %v", id, err)
		return err
	}

	hashedPassword, err := usermodel.HashPassword(password)
	if err != nil {
		logs.Errorf("生成hash密码失败, user ID: %d, error: %v", id, err)
		return err
	}

	result := db.Model(&user).Update("Password", hashedPassword)
	if result.Error != nil {
		logs.Errorf("更新用户密码失败, user ID: %d, error: %v", id, result.Error)
		return result.Error
	}

	return nil
}

// UpdateIsAdmin 修改用户权限
func UpdateIsAdmin(db *gorm.DB, isadmin bool, id uint) error {
	var user usermodel.User
	if err := db.First(&user, id).Error; err != nil {
		logs.Errorf("未找到用户: %d, error: %v", id, err)
		return err
	}

	result := db.Model(&user).Update("IsAdmin", isadmin)
	if result.Error != nil {
		logs.Errorf("更新用户权限失败, user ID: %d, error: %v", id, result.Error)
		return result.Error
	}

	return nil
}

// DeleteUserByID 根据ID删除用户
func DeleteUserByID(db *gorm.DB, id uint) error {
	result := db.Delete(&usermodel.User{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除用户失败, id: %d", id)
		return result.Error
	}
	logs.Infof("根据ID删除用户成功, id: %d", id)
	return result.Error
}

// GetUserByUserName 根据用户名查找用户
func GetUserByUserName(db *gorm.DB, username string) (*usermodel.User, error) {
	user := &usermodel.User{}
	result := db.Where("user_name = ?", username).First(user)
	if result.Error != nil {
		logs.Errorf("根据用户名查找用户失败, username: %s", username)
		return nil, result.Error
	}
	logs.Infof("根据用户名查找用户成功, username: %s", username)
	return user, nil
}

// GetAllUserInfo 获取所有的用户信息
func GetAllUserInfo(db *gorm.DB) ([]usermodel.User, error) {
	var users []usermodel.User
	result := db.Find(&users)
	if result.Error != nil {
		logs.Errorf("获取所有用户信息失败: %v", result.Error)
		return nil, result.Error
	}
	logs.Infof("获取所有用户信息成功")
	return users, nil
}
