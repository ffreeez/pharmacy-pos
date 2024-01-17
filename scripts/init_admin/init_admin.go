package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	"pharmacy-pos/pkg/db/models/user"
	"pharmacy-pos/pkg/util/logger"
)

func main() {

	logger.Init()
	config.Load()

	database, err := db.InitDB()
	if err != nil {
		panic("数据库初始化失败")
	}

	user := &usermodel.User{}
	user.ID = 1
	user.UserName = "admin"
	user.Password, _ = usermodel.HashPassword("password")
	user.IsAdmin = true
	database.Create(user)

}
