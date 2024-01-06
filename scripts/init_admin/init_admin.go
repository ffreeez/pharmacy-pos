package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	UserModel "pharmacy-pos/pkg/db/models"
	"pharmacy-pos/pkg/util/logger"
)

func main() {

	logger.Init()
	config.Load()

	database, err := db.InitDB()
	if err != nil {
		panic("数据库初始化失败")
	}

	user := &UserModel.User{}
	user.ID = 1
	user.UserName = "admin"
	user.Password, _ = UserModel.HashPassword("password")
	database.Create(user)

}
