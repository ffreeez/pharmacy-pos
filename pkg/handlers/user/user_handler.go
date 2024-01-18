package userhandler

import (
	"fmt"
	"strconv"

	usermodel "pharmacy-pos/pkg/db/models/user"
	"pharmacy-pos/pkg/middleware/jwt"
	userservice "pharmacy-pos/pkg/service/user"
	"pharmacy-pos/pkg/util/e"
	"pharmacy-pos/pkg/util/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler 处理用户相关的 HTTP 请求
type UserHandler struct {
	UserService *userservice.UserService
}

// NewUserHandler 创建一个新的 UserHandler 实例
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		UserService: userservice.NewUserService(db),
	}
}

// CreateUser 创建新用户
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user usermodel.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := uh.UserService.CreateUser(&user)
	if err != nil {
		if err.Error() == "用户名已存在" {
			response.Conflict(c, "Username already exists")
		} else {
			response.InternalServerError(c, "Failed to create user")
		}
		return
	}

	response.Created(c, user, "success")
}

// UpdateIsAdmin 修改用户权限
func (uh *UserHandler) UpdateIsAdmin(c *gin.Context) {
	var req struct {
		IsAdmin bool `json:"is_admin"`
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	userID := uint(id)
	err = uh.UserService.UpdateIsAdmin(req.IsAdmin, userID)
	if err != nil {
		response.InternalServerError(c, "Failed to update user")
		return
	}

	response.OK(c, gin.H{"is_admin": req.IsAdmin}, "success")
}

// ResetPassword 重设用户密码
func (uh *UserHandler) ResetPassword(c *gin.Context) {
	var req struct {
		Password string `json:"password"`
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	userID := uint(id)
	err = uh.UserService.ResetPassword(req.Password, userID)
	if err != nil {
		response.InternalServerError(c, "Failed to reset password")
		return
	}

	response.OK(c, nil, "Password reset successfully")
}

// DeleteUserByID 根据ID删除用户
func (uh *UserHandler) DeleteUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	userID := uint(id)
	err = uh.UserService.DeleteUserByID(userID)
	if err != nil {
		response.InternalServerError(c, "Failed to delete user")
		return
	}

	response.OK(c, gin.H{"message": "User deleted successfully"}, "success")
}

// Login 处理用户登录请求
func (uh *UserHandler) Login(c *gin.Context) {
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	user, err := uh.UserService.AuthenticateUser(loginInfo.Username, loginInfo.Password)
	if err != nil {
		response.Unauthorized(c, "Invalid username or password", e.CodeIllegalToken)
		return
	}

	tokenString, err := jwt.GenerateToken(user.UserName)
	if err != nil {
		response.Unauthorized(c, "Login failed, token generated fail", e.CodeIllegalToken)
		return
	}
	response.OK(c, gin.H{"token": tokenString}, "success")
}

// GetInfo 根据token获取用户信息
func (uh *UserHandler) GetInfo(c *gin.Context) {
	// 假设JWT中间件已经验证了token，并将userName存储在了Gin上下文中
	userName, exists := c.Get("username")
	if !exists {
		response.Unauthorized(c, "Failed to authenticate user", e.CodeIllegalToken)
		return
	}

	// 断言userName是一个字符串
	userNameStr, ok := userName.(string)
	if !ok {
		response.InternalServerError(c, "Server error occurred")
		return
	}

	// 使用用户名获取用户信息
	user, err := uh.UserService.GetUserByUserName(userNameStr)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(c, "User not found")
		} else {
			response.InternalServerError(c, "Failed to get user information")
		}
		return
	}

	// 返回用户信息
	response.OK(c, gin.H{"name": user.UserName, "avatar": user.Avatar}, "success")
}

// GetAllUserInfo 获取所有的用户信息，但只包含用户名、ID和是否是管理员
func (uh *UserHandler) GetAllUserInfo(c *gin.Context) {

	// SimplifiedUser 用于只返回必要的用户信息
	type SimplifiedUser struct {
		ID       uint   `json:"id"`
		UserName string `json:"username"`
		IsAdmin  bool   `json:"is_admin"`
	}
	// 假设 UserService 有一个方法 GetAllSimplifiedUsersInfo，只返回用户的ID、用户名和是否是管理员
	simplifiedUsers, err := uh.UserService.GetAllUserInfo()
	if err != nil {
		response.InternalServerError(c, "Failed to get all user info")
		return
	}

	// 创建一个用于响应的切片，只包含需要的字段
	var usersResponse []SimplifiedUser
	for _, user := range simplifiedUsers {
		usersResponse = append(usersResponse, SimplifiedUser{
			ID:       user.ID,
			UserName: user.UserName,
			IsAdmin:  user.IsAdmin, // 假设这是你模型中的字段
		})
	}

	// 如果成功获取到用户信息，返回OK状态码和用户列表
	response.OK(c, usersResponse, "success")
}

// GetUserByUserName 根据用户名获取用户信息
func (uh *UserHandler) GetUserByUserName(c *gin.Context) {
	// 从请求参数中获取用户名
	username := c.Param("username")
	if username == "" {
		response.BadRequest(c, "Username is required")
		return
	}

	fmt.Println(username)
	// 使用用户名获取用户信息
	user, err := uh.UserService.GetUserByUserName(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(c, "User not found")
		} else {
			response.InternalServerError(c, "Failed to get user information")
		}
		return
	}

	// SimplifiedUser 用于只返回必要的用户信息
	type SimplifiedUser struct {
		ID       uint   `json:"id"`
		UserName string `json:"username"`
		IsAdmin  bool   `json:"is_admin"`
	}

	// 创建一个用于响应的对象，只包含需要的字段
	userResponse := SimplifiedUser{
		ID:       user.ID,
		UserName: user.UserName,
		IsAdmin:  user.IsAdmin, // 假设这是你模型中的字段
	}

	// 如果成功获取到用户信息，返回OK状态码和用户信息
	response.OK(c, userResponse, "success")
}
