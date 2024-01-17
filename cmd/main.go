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
	user_protected_api.GET("/getbyusername/:username", userHandler.GetUserByUserName)
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

// 后端

// 添加分页
// 修改userupdate功能，user不存在返回失败
// 修改drugupdate功能，drug不存在返回失败
// 添加对中文的支持
// 添加校验方法中间件，验证用户输入
// 重构user相关代码，功能整合到service
// 添加用户头像链接
// 添加药品图片链接

// 前端
// 隐藏id
// 添加分页
// 删除框删除时需要确定
// 增加drug的筛选排序
// 增加drug的查询功能
// 添加药品图片
