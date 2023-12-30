package repository

import (
	"context"
	"pharmacy-pos/models"
	"time"

	"gorm.io/gorm"
)

// DiscountRepository 定义了操作折扣表的接口
type DiscountRepository interface {
	Create(ctx context.Context, discount *models.Discount) error
	Delete(ctx context.Context, discountID uint) error
	Update(ctx context.Context, discount *models.Discount) error
	FindByID(ctx context.Context, discountID uint) (*models.Discount, error)
	FindAll(ctx context.Context) ([]*models.Discount, error)
	FindActiveDiscounts(ctx context.Context, drugID uint, at time.Time) ([]*models.Discount, error)
}

// DiscountRepositoryImpl 实现了DiscountRepository接口
type DiscountRepositoryImpl struct {
	db *gorm.DB
}

// NewDiscountRepository 创建一个新的DiscountRepository实例
func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &DiscountRepositoryImpl{db}
}

// Create 添加一个新的折扣到数据库
func (repo *DiscountRepositoryImpl) Create(ctx context.Context, discount *models.Discount) error {
	return repo.db.WithContext(ctx).Create(discount).Error
}

// Delete 根据DiscountID删除一个折扣
func (repo *DiscountRepositoryImpl) Delete(ctx context.Context, discountID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Discount{}, discountID).Error
}

// Update 更新折扣信息
func (repo *DiscountRepositoryImpl) Update(ctx context.Context, discount *models.Discount) error {
	return repo.db.WithContext(ctx).Save(discount).Error
}

// FindByID 根据DiscountID查找一个折扣
func (repo *DiscountRepositoryImpl) FindByID(ctx context.Context, discountID uint) (*models.Discount, error) {
	var discount models.Discount
	err := repo.db.WithContext(ctx).First(&discount, discountID).Error
	return &discount, err
}

// FindAll 返回所有折扣的列表
func (repo *DiscountRepositoryImpl) FindAll(ctx context.Context) ([]*models.Discount, error) {
	var discounts []*models.Discount
	err := repo.db.WithContext(ctx).Find(&discounts).Error
	return discounts, err
}

// FindActiveDiscounts 返回在指定时间内有效的折扣列表
func (repo *DiscountRepositoryImpl) FindActiveDiscounts(ctx context.Context, drugID uint, at time.Time) ([]*models.Discount, error) {
	var discounts []*models.Discount
	err := repo.db.WithContext(ctx).Where("drug_id = ? AND start_date <= ? AND end_date >= ?", drugID, at, at).Find(&discounts).Error
	return discounts, err
}
