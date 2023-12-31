package models

import (
	"time"

	"gorm.io/gorm"
)

// Drug 药品表
type Drug struct {
	gorm.Model
	DrugName     string
	Description  string
	CategoryID   uint
	Category     Category `gorm:"foreignKey:CategoryID"`
	Price        float64  // 当前价格
	DiscountRate float64  // 当前折扣率
	Inventory    uint     // 库存数量
}

// Category 分类表
type Category struct {
	gorm.Model
	CategoryName string
	Drugs        []Drug `gorm:"foreignKey:CategoryID"`
}

// Sale 订单表
type Sale struct {
	gorm.Model
	UserID         uint
	User           User `gorm:"foreignKey:UserID"`
	MemberID       uint
	Member         Member       `gorm:"foreignKey:MemberID"`
	TotalAmount    float64      // 订单总金额
	DiscountAmount float64      // 订单折扣金额
	FinalAmount    float64      // 订单最终金额
	SaleDate       time.Time    // 销售日期
	SaleDetails    []SaleDetail `gorm:"foreignKey:SaleID"`
}

// SaleDetail 订单详情表
type SaleDetail struct {
	gorm.Model
	SaleID       uint
	Sale         Sale `gorm:"foreignKey:SaleID"`
	DrugID       uint
	Drug         Drug    `gorm:"foreignKey:DrugID"`
	Quantity     int     // 购买数量
	PricePerUnit float64 // 单价
}

// Member 会员表
type Member struct {
	gorm.Model
	Name        string
	PhoneNumber string
	Points      int // 会员积分
}

// User 用户表
type User struct {
	gorm.Model
	UserName string
	Password string
	IsAdmin  bool // 是否为管理员（店长）
}

// Log 操作日志表
type Log struct {
	gorm.Model
	UserID      uint
	User        User      `gorm:"foreignKey:UserID"`
	ActionType  string    // 操作类型
	Description string    // 操作描述
	ActionDate  time.Time // 操作日期
}
