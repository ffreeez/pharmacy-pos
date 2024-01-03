package user_handler

import (
	"strconv"

	"pharmacy-pos/pkg/db/models"
	"pharmacy-pos/pkg/service"
	"pharmacy-pos/pkg/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler 处理用户相关的 HTTP 请求
type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler 创建一个新的 UserHandler 实例
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		UserService: service.NewUserService(db),
	}
}

// GetUserByID 根据用户ID获取用户
func (uh *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.BadRequest(c, "Invalid user ID")
		return
	}

	userID := uint(id)
	user, err := uh.UserService.GetUserByID(userID)
	if err != nil {
		util.InternalServerError(c, "Failed to get user")
		return
	}

	util.OK(c, user)
}

// CreateUser 创建新用户
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.BadRequest(c, "Invalid input")
		return
	}

	err := uh.UserService.CreateUser(&user)
	if err != nil {
		util.InternalServerError(c, "Failed to create user")
		return
	}

	util.Created(c, user)
}

// UpdateUser 更新用户信息
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.BadRequest(c, "Invalid input")
		return
	}

	err := uh.UserService.UpdateUser(&user)
	if err != nil {
		util.InternalServerError(c, "Failed to update user")
		return
	}

	util.OK(c, user)
}

// DeleteUserByID 根据ID删除用户
func (uh *UserHandler) DeleteUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		util.BadRequest(c, "Invalid user ID")
		return
	}

	userID := uint(id)
	err = uh.UserService.DeleteUserByID(userID)
	if err != nil {
		util.InternalServerError(c, "Failed to delete user")
		return
	}

	util.OK(c, gin.H{"message": "User deleted successfully"})
}
