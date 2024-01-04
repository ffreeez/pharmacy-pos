package UserHandler

import (
	"strconv"

	UserModel "pharmacy-pos/pkg/db/models"
	jwt "pharmacy-pos/pkg/middleware"
	"pharmacy-pos/pkg/service"
	"pharmacy-pos/pkg/util/response"

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
		response.BadRequest(c, "Invalid user ID")
		return
	}

	userID := uint(id)
	user, err := uh.UserService.GetUserByID(userID)
	if err != nil {
		response.InternalServerError(c, "Failed to get user")
		return
	}

	response.OK(c, user)
}

// CreateUser 创建新用户
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user UserModel.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err := uh.UserService.CreateUser(&user)
    if err != nil {
        if err.Error() == "用户名已存在" {
            response.InternalServerError(c, "Username already exists")
        } else {
            response.InternalServerError(c, "Failed to create user")
        }
        return
    }

	response.Created(c, user)
}

// UpdateUser 更新用户信息
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}

	var user UserModel.User
	userID := uint(id)
	user.ID = userID
	println(user.ID)

	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, "Invalid input")
		return
	}

	err = uh.UserService.UpdateUser(&user)
	if err != nil {
		response.InternalServerError(c, "Failed to update user")
		return
	}

	response.OK(c, user)
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

	response.OK(c, gin.H{"message": "User deleted successfully"})
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
		response.Unauthorized(c, "Invalid username or password")
		return
	}

	tokenString, err := jwt.GenerateToken(user.UserName)
	if err != nil {
		response.Unauthorized(c, "Login failed, token generated fail")
		return
	}
	response.OK(c, gin.H{"token": tokenString})
}
