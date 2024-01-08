package response

import (
	"net/http"
	"pharmacy-pos/pkg/util/e"

	"github.com/gin-gonic/gin"
)

// OK 发送 HTTP 200 OK 响应
func OK(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, gin.H{"code": e.CodeSuccess, "data": data, "message": message})
}

// Created 发送 HTTP 201 Created 响应
func Created(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusCreated, gin.H{"code": e.CodeSuccess, "data": data, "message": message})
}

// BadRequest 发送 HTTP 400 Bad Request 响应
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": message})
}

// InternalServerError 发送 HTTP 500 Internal Server Error 响应
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": message})
}

// Unauthorized 发送 401 未授权的响应
func Unauthorized(c *gin.Context, message string, code int) {
	// 使用额外的 code 参数来区分不同的授权错误
	c.JSON(http.StatusUnauthorized, gin.H{"code": code, "message": message})
}

// Conflict 发送 HTTP 409 Conflict 响应
func Conflict(c *gin.Context, message string) {
	c.JSON(http.StatusConflict, gin.H{"code": http.StatusConflict, "message": message})
}

// NotFound 发送 HTTP 404 NotFound 响应
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": message})
}
