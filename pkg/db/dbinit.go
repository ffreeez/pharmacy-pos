package db

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db/models"
	logger "pharmacy-pos/pkg/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() {

	logs := logger.GetLogger()

	config.Load()
	dsn := config.GetDb()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Error("数据库连接失败")
	}

	logs.Info("数据库连接成功")

	DB.AutoMigrate(&models.User{})
}
