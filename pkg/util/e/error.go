package e

import (
	"errors"
)

var ErrUsernameExists = errors.New("用户名已存在")
var ErrNotFound = errors.New("对象未找到")

const (
	CodeSuccess      = 20000
	CodeIllegalToken = 50008 // 非法token
	CodeOtherClient  = 50012 // 其他客户端登录
	CodeTokenExpired = 50014 // Token过期

)
