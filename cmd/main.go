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

	router.GET("/users/:id", userHandler.GetUserByID)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUserByID)

	router.Run(":8080")
}
