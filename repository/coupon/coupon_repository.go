package repository

import (
	"context"
	"pharmacy-pos/models"
	"time"

	"gorm.io/gorm"
)

// CouponRepository 定义了操作优惠券表的接口
type CouponRepository interface {
	Create(ctx context.Context, coupon *models.Coupon) error
	Delete(ctx context.Context, couponID uint) error
	Update(ctx context.Context, coupon *models.Coupon) error
	FindByID(ctx context.Context, couponID uint) (*models.Coupon, error)
	FindAll(ctx context.Context) ([]*models.Coupon, error)
	FindActiveCoupons(ctx context.Context, at time.Time) ([]*models.Coupon, error)
}

// CouponRepositoryImpl 实现了CouponRepository接口
type CouponRepositoryImpl struct {
	db *gorm.DB
}

// NewCouponRepository 创建一个新的CouponRepository实例
func NewCouponRepository(db *gorm.DB) CouponRepository {
	return &CouponRepositoryImpl{db}
}

// Create 添加一个新的优惠券到数据库
func (repo *CouponRepositoryImpl) Create(ctx context.Context, coupon *models.Coupon) error {
	return repo.db.WithContext(ctx).Create(coupon).Error
}

// Delete 根据CouponID删除一个优惠券
func (repo *CouponRepositoryImpl) Delete(ctx context.Context, couponID uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Coupon{}, couponID).Error
}

// Update 更新优惠券信息
func (repo *CouponRepositoryImpl) Update(ctx context.Context, coupon *models.Coupon) error {
	return repo.db.WithContext(ctx).Save(coupon).Error
}

// FindByID 根据CouponID查找一个优惠券
func (repo *CouponRepositoryImpl) FindByID(ctx context.Context, couponID uint) (*models.Coupon, error) {
	var coupon models.Coupon
	err := repo.db.WithContext(ctx).First(&coupon, couponID).Error
	return &coupon, err
}

// FindAll 返回所有优惠券的列表
func (repo *CouponRepositoryImpl) FindAll(ctx context.Context) ([]*models.Coupon, error) {
	var coupons []*models.Coupon
	err := repo.db.WithContext(ctx).Find(&coupons).Error
	return coupons, err
}

// FindActiveCoupons 返回在指定时间内有效的优惠券列表
func (repo *CouponRepositoryImpl) FindActiveCoupons(ctx context.Context, at time.Time) ([]*models.Coupon, error) {
	var coupons []*models.Coupon
	err := repo.db.WithContext(ctx).Where("start_date <= ? AND end_date >= ?", at, at).Find(&coupons).Error
	return coupons, err
}
