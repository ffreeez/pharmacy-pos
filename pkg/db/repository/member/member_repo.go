package memberrepo

import (
	"pharmacy-pos/pkg/db/models/member"
	logger "pharmacy-pos/pkg/util/logger"
	"time"

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

// AddCouponToMember 为会员添加优惠券
func AddCouponToMember(db *gorm.DB, coupon *membermodel.Coupon) error {
	result := db.Create(coupon)
	if result.Error != nil {
		logs.Errorf("为会员添加优惠券失败, MemberID: %d, error: %v", coupon.MemberID, result.Error)
		return result.Error
	}
	logs.Infof("为会员添加优惠券成功, MemberID: %d", coupon.MemberID)
	return nil
}

// UseCoupon 使用优惠券
func UseCoupon(db *gorm.DB, couponID uint) error {
	result := db.Model(&membermodel.Coupon{}).Where("id = ? AND used = ?", couponID, false).Updates(map[string]interface{}{"Used": true, "ValidUntil": time.Now()})
	if result.Error != nil {
		logs.Errorf("使用优惠券失败, CouponID: %d, error: %v", couponID, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logs.Errorf("优惠券不存在或已被使用, CouponID: %d", couponID)
		return gorm.ErrRecordNotFound
	}
	logs.Infof("使用优惠券成功, CouponID: %d", couponID)
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

// GetCouponsByMemberID 获取会员的所有优惠券
func GetCouponsByMemberID(db *gorm.DB, memberID uint) ([]membermodel.Coupon, error) {
	var coupons []membermodel.Coupon
	result := db.Where("member_id = ?", memberID).Find(&coupons)
	if result.Error != nil {
		logs.Errorf("获取会员的所有优惠券失败, MemberID: %d, error: %v", memberID, result.Error)
		return nil, result.Error
	}
	logs.Infof("获取会员的所有优惠券成功, MemberID: %d", memberID)
	return coupons, nil
}
