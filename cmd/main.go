package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	drughandler "pharmacy-pos/pkg/handlers/drug"
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
	jwt.InitJWTKey()

	userHandler := userhandler.NewUserHandler(database)
	drugHandler := drughandler.NewDrugHandler(database)

	router.POST("/login", userHandler.Login)

	user_protected_api := router.Group("/users")
	user_protected_api.Use(jwt.JWTAuthMiddleware())
	user_protected_api.POST("/create", userHandler.CreateUser)
	user_protected_api.GET("/get/:id", userHandler.GetUserByID)
	user_protected_api.PUT("/update/password/:id", userHandler.ResetPassword)
	user_protected_api.PUT("/update/isadmin/:id", userHandler.UpdateIsAdmin)
	user_protected_api.DELETE("/delete/:id", userHandler.DeleteUserByID)
	user_protected_api.GET("/getall", userHandler.GetAllUserInfo)
	user_protected_api.GET("/info", userHandler.GetInfo)

	drug_protected_api := router.Group("/drugs")
	drug_protected_api.Use(jwt.JWTAuthMiddleware())
	drug_protected_api.POST("/create", drugHandler.CreateDrug)
	drug_protected_api.GET("/get/:id", drugHandler.GetDrugByID)
	drug_protected_api.PUT("/update/:id", drugHandler.UpdateDrug)
	drug_protected_api.DELETE("/delete/:id", drugHandler.DeleteDrugByID)
	drug_protected_api.GET("/getall", drugHandler.GetAllDrugs)

	category_protected_api := router.Group("/categories")
	category_protected_api.Use(jwt.JWTAuthMiddleware())
	category_protected_api.POST("/create", drugHandler.CreateCategory)
	category_protected_api.GET("/get/:id", drugHandler.GetCategoryByID)
	category_protected_api.PUT("/update/:id", drugHandler.UpdateCategory)
	category_protected_api.DELETE("/delete/:id", drugHandler.DeleteCategoryByID)
	category_protected_api.GET("/getall", drugHandler.GetAllCategorys)

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

// TODO
// 重构Users相关内容，将handler中的不合理结构转移到service
// 完成所有数据模型的repo
// 整理数据模型文档，确定数据模型
// 整理api文档，确定最基础的api
// 完成所有数据模型的service
// 完成所有数据模型的handler
// 使用postman进行所有的api测试
