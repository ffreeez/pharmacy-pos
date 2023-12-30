package repository

import (
	"context"
	"pharmacy-pos/models"

	"gorm.io/gorm"
)

// SaleRepository 定义了操作订单表的接口
type SaleRepository interface {
	Create(ctx context.Context, sale *models.Sale) error
	Delete(ctx context.Context, saleID uint) error
	Update(ctx context.Context, sale *models.Sale) error
	FindByID(ctx context.Context, saleID uint) (*models.Sale, error)
	FindAll(ctx context.Context) ([]*models.Sale, error)
}

// SaleRepositoryImpl 实现了SaleRepository接口
type SaleRepositoryImpl struct {
	db *gorm.DB
}

// NewSaleRepository 创建一个新的SaleRepository实例
func NewSaleRepository(db *gorm.DB) SaleRepository {
	return &SaleRepositoryImpl{db}
}

// Create 添加一个新的订单到数据库
func (repo *SaleRepositoryImpl) Create(ctx context.Context, sale *models.Sale) error {
	return repo.db.WithContext(ctx).Create(sale).Error
}

// Delete 根据SaleID删除一个订单
func (repo *SaleRepositoryImpl) Delete(ctx context.Context, saleID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Sale{}, saleID).Error
}

// Update 更新订单信息
func (repo *SaleRepositoryImpl) Update(ctx context.Context, sale *models.Sale) error {
	return repo.db.WithContext(ctx).Save(sale).Error
}

// FindByID 根据SaleID查找一个订单
func (repo *SaleRepositoryImpl) FindByID(ctx context.Context, saleID uint) (*models.Sale, error) {
	var sale models.Sale
	err := repo.db.WithContext(ctx).First(&sale, saleID).Error
	return &sale, err
}

// FindAll 返回所有订单的列表
func (repo *SaleRepositoryImpl) FindAll(ctx context.Context) ([]*models.Sale, error) {
	var sales []*models.Sale
	err := repo.db.WithContext(ctx).Find(&sales).Error
	return sales, err
}
