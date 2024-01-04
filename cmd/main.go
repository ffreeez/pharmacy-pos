package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	UserHandler "pharmacy-pos/pkg/handlers"
	jwt "pharmacy-pos/pkg/middleware"
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

	router.POST("/login", userHandler.Login)

	jwt.InitJWTKey()
	protected := router.Group("/users")
	protected.Use(jwt.JWTAuthMiddleware())

	protected.POST("/create", userHandler.CreateUser)
	protected.GET("/get/:id", userHandler.GetUserByID)
	protected.PUT("/update/:id", userHandler.UpdateUser)
	protected.DELETE("/delete/:id", userHandler.DeleteUserByID)

	router.Run(":8080")
}

// TODO
// 修改http返回内容
// 用户名唯一检查
