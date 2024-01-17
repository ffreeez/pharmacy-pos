package drugmodel

import (
	"gorm.io/gorm"
)

// Category 表示药品的分类
type Category struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Drugs []Drug `gorm:"foreignKey:CategoryID"`
}

// Drug 表示药品表
type Drug struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text;not null"`
	CategoryID  uint      // 外键 (属于 - Category)
	Category    Category  // 关联
	Price       Price     // 价格信息 (hasOne - Price)
	Inventory   Inventory // 库存信息 (hasOne - Inventory)
}

// Price 表示药品价格
type Price struct {
	gorm.Model
	DrugID uint    `gorm:"uniqueIndex"` // 药品ID
	Value  float64 `gorm:"not null"`    // 价格值
}

// Inventory 表示药品库存
type Inventory struct {
	gorm.Model
	DrugID uint `gorm:"uniqueIndex"` // 药品ID
	Amount int  `gorm:"not null"`    // 库存数量
}
