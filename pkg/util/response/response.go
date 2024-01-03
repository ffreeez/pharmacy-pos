package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OK 发送 HTTP 200 OK 响应
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Created 发送 HTTP 201 Created 响应
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"data": data})
}

// BadRequest 发送 HTTP 400 Bad Request 响应
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}

// InternalServerError 发送 HTTP 500 Internal Server Error 响应
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": message})
}
