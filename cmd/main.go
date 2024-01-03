package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	"pharmacy-pos/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	config.Load()
	router := gin.Default()
	router.GET("/users/:id", user_handler..GetUserByID)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUserByID)
	router.Run(":8080")
}
