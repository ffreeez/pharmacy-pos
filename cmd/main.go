package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	userhandler "pharmacy-pos/pkg/handlers/user"
	"pharmacy-pos/pkg/middleware/cors"
	"pharmacy-pos/pkg/middleware/jwt"
	"pharmacy-pos/pkg/util/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(database *gorm.DB) *gin.Engine {

	router := gin.Default()
	router.Use(cors.Cors())
	userHandler := userhandler.NewUserHandler(database)

	router.POST("/login", userHandler.Login)

	jwt.InitJWTKey()
	protected := router.Group("/users")
	protected.Use(jwt.JWTAuthMiddleware())
	protected.POST("/create", userHandler.CreateUser)
	protected.GET("/get/:id", userHandler.GetUserByID)
	protected.PUT("/update/password/:id", userHandler.ResetPassword)
	protected.PUT("/update/isadmin/:id", userHandler.UpdateIsAdmin)
	protected.DELETE("/delete/:id", userHandler.DeleteUserByID)
	protected.GET("/getall", userHandler.GetAllUserInfo)
	protected.GET("/info", userHandler.GetInfo)

	return router
}

func main() {

	logger.Init()
	database, err := db.InitDB()
	if err != nil {
		panic("数据库初始化失败")
	}
	config.Load()

	router := setupRouter(database)

	router.Run(":8080")
}
