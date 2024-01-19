package memberrepo

import (
	membermodel "pharmacy-pos/pkg/db/models/member"
	logger "pharmacy-pos/pkg/util/logger"

	"gorm.io/gorm"
)

var logs = logger.GetLogger()

// CreateMember 创建新会员
func CreateMember(db *gorm.DB, member *membermodel.Member) error {
	result := db.Create(member)
	if result.Error != nil {
		logs.Errorf("创建新会员失败, phone: %s, error: %v", member.Phone, result.Error)
		return result.Error
	}
	logs.Infof("创建新会员成功, phone: %s", member.Phone)
	return nil
}

// GetMemberByID 根据会员ID获取会员信息
func GetMemberByID(db *gorm.DB, id uint) (*membermodel.Member, error) {
	member := &membermodel.Member{}
	result := db.Preload("Coupons").First(member, id)
	if result.Error != nil {
		logs.Errorf("根据会员ID获取会员信息失败, ID: %d, error: %v", id, result.Error)
		return nil, result.Error
	}
	logs.Infof("根据会员ID获取会员信息成功, ID: %d", id)
	return member, nil
}

// UpdateMember 更新会员信息
func UpdateMember(db *gorm.DB, member *membermodel.Member) error {
	result := db.Save(member)
	if result.Error != nil {
		logs.Errorf("更新会员信息失败, ID: %d, error: %v", member.ID, result.Error)
		return result.Error
	}
	logs.Infof("更新会员信息成功, ID: %d", member.ID)
	return nil
}

// DeleteMemberByID 根据ID删除会员
func DeleteMemberByID(db *gorm.DB, id uint) error {
	result := db.Delete(&membermodel.Member{}, id)
	if result.Error != nil {
		logs.Errorf("根据ID删除会员失败, ID: %d, error: %v", id, result.Error)
		return result.Error
	}
	logs.Infof("根据ID删除会员成功, ID: %d", id)
	return nil
}

// GetAllMembers 获取所有会员信息
func GetAllMembers(db *gorm.DB) ([]membermodel.Member, error) {
	var members []membermodel.Member
	result := db.Preload("Coupons").Find(&members)
	if result.Error != nil {
		logs.Errorf("获取所有会员信息失败: %v", result.Error)
		return nil, result.Error
	}
	logs.Infof("获取所有会员信息成功")
	return members, nil
}

// CreateCoupon 创建优惠券
func CreateCoupon(db *gorm.DB, coupon *membermodel.Coupon) error {
	result := db.Create(coupon)
	if result.Error != nil {
		logs.Errorf("添加优惠券失败, error: %v", result.Error)
		return result.Error
	}
	logs.Info("添加优惠券成功")
	return nil
}

// GetCouponByID 根据id查找优惠券
func GetCouponByID(db *gorm.DB, couponID uint) (*membermodel.Coupon, error) {
	var coupon membermodel.Coupon
	result := db.First(&coupon, couponID)
	if result.Error != nil {
		logs.Errorf("查找优惠券失败, CouponID: %d, error: %v", couponID, result.Error)
		return nil, result.Error
	}
	logs.Infof("查找优惠券成功, CouponID: %d", couponID)
	return &coupon, nil
}

// UpdateCoupon 更新优惠券信息
func UpdateCoupon(db *gorm.DB, coupon *membermodel.Coupon) error {
	result := db.Save(coupon)
	if result.Error != nil {
		logs.Errorf("更新优惠券失败, CouponID: %d, error: %v", coupon.ID, result.Error)
		return result.Error
	}
	logs.Infof("更新优惠券成功, CouponID: %d", coupon.ID)
	return nil
}

// DeleteCoupon 删除优惠券
func DeleteCoupon(db *gorm.DB, couponID uint) error {
	result := db.Delete(&membermodel.Coupon{}, couponID)
	if result.Error != nil {
		logs.Errorf("删除优惠券失败, CouponID: %d, error: %v", couponID, result.Error)
		return result.Error
	}
	logs.Infof("删除优惠券成功, CouponID: %d", couponID)
	return nil
}

// GetAllCoupons 获取所有优惠券
func GetAllCoupons(db *gorm.DB) ([]membermodel.Coupon, error) {
	var coupons []membermodel.Coupon
	result := db.Find(&coupons)
	if result.Error != nil {
		logs.Errorf("获取所有优惠券失败, error: %v", result.Error)
		return nil, result.Error
	}
	logs.Info("获取所有优惠券成功")
	return coupons, nil
}
