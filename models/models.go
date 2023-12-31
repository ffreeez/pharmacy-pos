package models

import (
	"time"

	"gorm.io/gorm"
)

// 用户表（Users）数据模型
type User struct {
	gorm.Model
	UserID   uint `gorm:"primaryKey"`
	UserName string
	Password string
	Role     string
}

// 药品表（Drugs）数据模型
type Drug struct {
	gorm.Model
	DrugID        uint `gorm:"primaryKey"`
	Name          string
	Category      string
	Price         float64
	Description   string
	StockQuantity int
}

// 价格表（Prices）数据模型
type Price struct {
	gorm.Model
	PriceID uint `gorm:"primaryKey"`
	DrugID  uint `gorm:"foreignKey:DrugID"`
	Price   float64
}

// 折扣表（Discounts）数据模型
type Discount struct {
	gorm.Model
	DiscountID   uint `gorm:"primaryKey"`
	DrugID       uint `gorm:"foreignKey:DrugID"`
	DiscountRate float64
	StartDate    time.Time
	EndDate      time.Time
}

// 优惠券表（Coupons）数据模型
type Coupon struct {
	gorm.Model
	CouponID          uint `gorm:"primaryKey"`
	Type              string
	DiscountValue     float64
	MinPurchaseAmount float64
	StartDate         time.Time
	EndDate           time.Time
}

// 会员表（Members）数据模型
type Member struct {
	gorm.Model
	MemberID    uint `gorm:"primaryKey"`
	Name        string
	PhoneNumber string
	Points      int
	JoinDate    time.Time
}

// 订单表（Sales）数据模型
type Sale struct {
	gorm.Model
	SaleID         uint `gorm:"primaryKey"`
	UserID         uint
	User           User `gorm:"foreignKey:UserID"`
	MemberID       uint
	Member         Member `gorm:"foreignKey:MemberID"`
	TotalAmount    float64
	DiscountAmount float64
	FinalAmount    float64
	SaleDate       time.Time
}

// 订单详情表（SaleDetails）数据模型
type SaleDetail struct {
	gorm.Model
	SaleDetailID uint `gorm:"primaryKey"`
	SaleID       uint
	Sale         Sale `gorm:"foreignKey:SaleID"`
	DrugID       uint
	Drug         Drug `gorm:"foreignKey:DrugID"`
	Quantity     int
	PricePerUnit float64
}

// 库存表（Inventory）数据模型
type Inventory struct {
	gorm.Model
	InventoryID uint `gorm:"primaryKey"`
	DrugID      uint
	Drug        Drug `gorm:"foreignKey:DrugID"`
	Quantity    int
	LastUpdated time.Time
}

// 操作日志表（Logs）数据模型
type Log struct {
	gorm.Model
	LogID       uint `gorm:"primaryKey"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	ActionType  string
	Description string
	ActionDate  time.Time
}
