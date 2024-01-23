package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
	drughandler "pharmacy-pos/pkg/handlers/drug"
	memberhandler "pharmacy-pos/pkg/handlers/member"
	orderhandler "pharmacy-pos/pkg/handlers/order"
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
	memberHandler := memberhandler.NewMemberHandler(database)
	orderHandler := orderhandler.NewOrderHandler(database)

	router.POST("/login", userHandler.Login)

	user_protected_api := router.Group("/users")
	user_protected_api.Use(jwt.JWTAuthMiddleware())
	user_protected_api.POST("/create", userHandler.CreateUser)
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
	drug_protected_api.GET("/getbydrugname/:drugname", drugHandler.GetDrugByDrugName)

	category_protected_api := router.Group("/categories")
	category_protected_api.Use(jwt.JWTAuthMiddleware())
	category_protected_api.POST("/create", drugHandler.CreateCategory)
	category_protected_api.GET("/get/:id", drugHandler.GetCategoryByID)
	category_protected_api.PUT("/update/:id", drugHandler.UpdateCategory)
	category_protected_api.DELETE("/delete/:id", drugHandler.DeleteCategoryByID)
	category_protected_api.GET("/getall", drugHandler.GetAllCategories)
	category_protected_api.GET("/getbycategoryname/:name", drugHandler.GetCategoryByName)

	member_protected_api := router.Group("/members")
	member_protected_api.Use(jwt.JWTAuthMiddleware())
	member_protected_api.POST("/create", memberHandler.CreateMember)
	member_protected_api.GET("/get/:id", memberHandler.GetMemberByID)
	member_protected_api.PUT("/update/:id", memberHandler.UpdateMember)
	member_protected_api.DELETE("/delete/:id", memberHandler.DeleteMemberByID)
	member_protected_api.GET("/getall", memberHandler.GetAllMembers)

	coupon_protected_api := router.Group("/coupons")
	coupon_protected_api.Use(jwt.JWTAuthMiddleware())
	coupon_protected_api.POST("/create", memberHandler.CreateMember)
	coupon_protected_api.GET("/get/:id", memberHandler.GetMemberByID)
	coupon_protected_api.PUT("/update/:id", memberHandler.UpdateMember)
	coupon_protected_api.DELETE("/delete/:id", memberHandler.DeleteMemberByID)
	coupon_protected_api.GET("/getall", memberHandler.GetAllMembers)

	order_protected_api := router.Group("/orders")
	order_protected_api.Use(jwt.JWTAuthMiddleware())
	order_protected_api.POST("/create", orderHandler.CreateOrder)
	order_protected_api.GET("/get/:id", orderHandler.GetOrderByID)
	order_protected_api.PUT("/update/:id", orderHandler.UpdateOrder)
	order_protected_api.DELETE("/delete/:id", orderHandler.DeleteOrderByID)
	order_protected_api.GET("/getall", orderHandler.GetAllOrders)

	orderitem_protected_api := router.Group("/orderitems")
	orderitem_protected_api.Use(jwt.JWTAuthMiddleware())
	orderitem_protected_api.POST("/create", orderHandler.CreateOrderItem)
	orderitem_protected_api.GET("/get/:id", orderHandler.GetOrderItemByID)
	orderitem_protected_api.PUT("/update/:id", orderHandler.UpdateOrderItem)
	orderitem_protected_api.DELETE("/delete/:id", orderHandler.DeleteOrderItemByID)
	orderitem_protected_api.GET("/getall", orderHandler.GetAllOrderItems)

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

	router.Run(config.AppConfig.Service.ServerPort)
}
