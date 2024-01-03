package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	UserHandler "pharmacy-pos/pkg/handlers"
	"pharmacy-pos/pkg/util/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	logger.Init()
	database, err := db.InitDB()
	if err != nil {
		panic("数据库初始化失败")
	}
	config.Load()

	router := gin.Default()

	userHandler := UserHandler.NewUserHandler(database)

	router.POST("/register", userHandler.CreateUser)
	router.POST("/login", userHandler.Login)
	
	router.GET("/users/:id", userHandler.GetUserByID)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUserByID)

	router.Run(":8080")
}
// TODO
// 修改http返回内容
// 添加登录令牌
// 测试用户的增删改查
// 为用户的增删改查添加登陆令牌和权限控制