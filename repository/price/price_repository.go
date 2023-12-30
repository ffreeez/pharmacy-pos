package repository

import (
	"context"
	"pharmacy-pos/models"

	"gorm.io/gorm"
)

// PriceRepository 定义了操作价格表的接口
type PriceRepository interface {
	Create(ctx context.Context, price *models.Price) error
	Delete(ctx context.Context, priceID uint) error
	Update(ctx context.Context, price *models.Price) error
	FindByID(ctx context.Context, priceID uint) (*models.Price, error)
	FindAll(ctx context.Context) ([]*models.Price, error)
}

// PriceRepositoryImpl 实现了PriceRepository接口
type PriceRepositoryImpl struct {
	db *gorm.DB
}

// NewPriceRepository 创建一个新的PriceRepository实例
func NewPriceRepository(db *gorm.DB) PriceRepository {
	return &PriceRepositoryImpl{db}
}

// Create 添加一个新的价格到数据库
func (repo *PriceRepositoryImpl) Create(ctx context.Context, price *models.Price) error {
	return repo.db.WithContext(ctx).Create(price).Error
}

// Delete 根据PriceID删除一个价格
func (repo *PriceRepositoryImpl) Delete(ctx context.Context, priceID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Price{}, priceID).Error
}

// Update 更新价格信息
func (repo *PriceRepositoryImpl) Update(ctx context.Context, price *models.Price) error {
	return repo.db.WithContext(ctx).Save(price).Error
}

// FindByID 根据PriceID查找一个价格
func (repo *PriceRepositoryImpl) FindByID(ctx context.Context, priceID uint) (*models.Price, error) {
	var price models.Price
	err := repo.db.WithContext(ctx).First(&price, priceID).Error
	return &price, err
}

// FindAll 返回所有价格的列表
func (repo *PriceRepositoryImpl) FindAll(ctx context.Context) ([]*models.Price, error) {
	var prices []*models.Price
	err := repo.db.WithContext(ctx).Find(&prices).Error
	return prices, err
}
