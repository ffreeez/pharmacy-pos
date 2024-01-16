package membermodel

import (
	"time"

	"gorm.io/gorm"
)

type CouponTypeCode uint

const (
	DiscountType      CouponTypeCode = iota // 折扣类型
	FullReductionType                       // 满减类型
)

// Member 会员表
type Member struct {
	gorm.Model
	Name    string   `gorm:"type:varchar(100);not null"` // 姓名
	Phone   string   `gorm:"type:varchar(100);unique"`   // 电话，假设电话号码是唯一的
	Point   uint     // 积分
	Coupons []Coupon `gorm:"foreignKey:MemberID"` // 会员的优惠券
}

// CouponType 优惠券类型表
type CouponType struct {
	gorm.Model
	Type           CouponTypeCode // 优惠券类型，使用自定义枚举
	Description    string         // 描述
	MinAmount      float64        // 最低消费金额
	DiscountAmount float64        // 减免金额或折扣比例
	Coupons        []Coupon       // 关联的优惠券
}

// Coupon 优惠券表
type Coupon struct {
	gorm.Model
	CouponTypeID uint      // 优惠券类型ID
	MemberID     uint      // 会员ID
	ValidUntil   time.Time // 有效期
	Used         bool      // 是否已使用
}
