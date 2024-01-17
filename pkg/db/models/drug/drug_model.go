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
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text;not null"`
	CategoryID  uint   // 外键 (属于 - Category)
	Category    Category
	Price       float64
	Inventory   uint
}
