package syslogmodel

import (
	"gorm.io/gorm"
)

type Syslog struct {
	gorm.Model
	UserID   uint   `gorm:"index" json:"userId"`     // 关联的用户ID，用于索引以提高查询效率
	UserName string `json:"username"`                // 用户名，记录执行操作的用户
	Action   string `gorm:"type:text" json:"action"` // 具体的操作描述
}
