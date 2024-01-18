package drugmodel

import (
	"gorm.io/gorm"
)

// Category 表示药品的分类
type Category struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null; charset:utf8mb4"`
	Drugs []Drug `gorm:"foreignKey:CategoryID"`
}

// Drug 表示药品表
type Drug struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null; charset:utf8mb4"`
	Description string `gorm:"type:varchar(100);not null; charset:utf8mb4"`
	CategoryID  uint   `json:"category_id"`
	Category    Category
	Price       float64
	Inventory   uint
	ImageURL    string
}
