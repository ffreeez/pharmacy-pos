package repository

import (
	"context"
	"pharmacy-pos/models"

	"gorm.io/gorm"
)

// SaleDetailRepository 定义了操作订单详情表的接口
type SaleDetailRepository interface {
	Create(ctx context.Context, saleDetail *models.SaleDetail) error
	Delete(ctx context.Context, saleDetailID uint) error
	Update(ctx context.Context, saleDetail *models.SaleDetail) error
	FindByID(ctx context.Context, saleDetailID uint) (*models.SaleDetail, error)
	FindAll(ctx context.Context) ([]*models.SaleDetail, error)
	FindBySaleID(ctx context.Context, saleID uint) ([]*models.SaleDetail, error)
}

// SaleDetailRepositoryImpl 实现了SaleDetailRepository接口
type SaleDetailRepositoryImpl struct {
	db *gorm.DB
}

// NewSaleDetailRepository 创建一个新的SaleDetailRepository实例
func NewSaleDetailRepository(db *gorm.DB) SaleDetailRepository {
	return &SaleDetailRepositoryImpl{db}
}

// Create 添加一个新的订单详情记录到数据库
func (repo *SaleDetailRepositoryImpl) Create(ctx context.Context, saleDetail *models.SaleDetail) error {
	return repo.db.WithContext(ctx).Create(saleDetail).Error
}

// Delete 根据SaleDetailID删除一个订单详情记录
func (repo *SaleDetailRepositoryImpl) Delete(ctx context.Context, saleDetailID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.SaleDetail{}, saleDetailID).Error
}

// Update 更新订单详情信息
func (repo *SaleDetailRepositoryImpl) Update(ctx context.Context, saleDetail *models.SaleDetail) error {
	return repo.db.WithContext(ctx).Save(saleDetail).Error
}

// FindByID 根据SaleDetailID查找一个订单详情记录
func (repo *SaleDetailRepositoryImpl) FindByID(ctx context.Context, saleDetailID uint) (*models.SaleDetail, error) {
	var saleDetail models.SaleDetail
	err := repo.db.WithContext(ctx).First(&saleDetail, saleDetailID).Error
	return &saleDetail, err
}

// FindAll 返回所有订单详情记录的列表
func (repo *SaleDetailRepositoryImpl) FindAll(ctx context.Context) ([]*models.SaleDetail, error) {
	var saleDetails []*models.SaleDetail
	err := repo.db.WithContext(ctx).Find(&saleDetails).Error
	return saleDetails, err
}

// FindBySaleID 根据SaleID查找所有相关的订单详情记录
func (repo *SaleDetailRepositoryImpl) FindBySaleID(ctx context.Context, saleID uint) ([]*models.SaleDetail, error) {
	var saleDetails []*models.SaleDetail
	err := repo.db.WithContext(ctx).Where("sale_id = ?", saleID).Find(&saleDetails).Error
	return saleDetails, err
}
