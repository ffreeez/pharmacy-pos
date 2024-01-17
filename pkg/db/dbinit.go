package db

import (
	"pharmacy-pos/pkg/config"
	drugmodel "pharmacy-pos/pkg/db/models/drug"
	usermodel "pharmacy-pos/pkg/db/models/user"
	logger "pharmacy-pos/pkg/util/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() (database *gorm.DB, err error) {

	logs := logger.GetLogger()

	config.Load()
	dsn := config.GetDb()

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Error("数据库连接失败")
		return nil, err
	}

	DB.AutoMigrate(&usermodel.User{})
	DB.AutoMigrate(&drugmodel.Drug{})
	DB.AutoMigrate(&drugmodel.Category{})
	logs.Info("数据库连接成功")

	return DB, nil
}
